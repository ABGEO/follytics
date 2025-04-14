package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/go-github/v68/github"
	"github.com/google/uuid"

	"github.com/abgeo/follytics/internal/helper"
	"github.com/abgeo/follytics/internal/model"
	"github.com/abgeo/follytics/internal/query"
	"github.com/abgeo/follytics/internal/repository"
)

type UserService interface {
	Me(ctx context.Context) (*model.User, error)
	Sync(ctx context.Context) (*model.User, error)
	GetRegularUsers(ctx context.Context, querier query.Querier) ([]*model.User, error)
	StoreGitHubFollowers(ctx context.Context, user *model.User, followers []*github.User) error
	GetFollowers(ctx context.Context, userID uuid.UUID, querier query.Querier) ([]*model.User, error)
	GetFollowEvents(ctx context.Context, userID uuid.UUID, querier query.Querier) ([]*model.Event, error)
}

type User struct {
	logger    *slog.Logger
	txManager *repository.TransactionManager

	eventRepo repository.EventRepository
	userRepo  repository.UserRepository

	authSvc   AuthService
	githubSvc GithubService
}

var _ UserService = (*User)(nil)

func NewUser(
	logger *slog.Logger,
	txManager *repository.TransactionManager,
	eventRepo repository.EventRepository,
	userRepo repository.UserRepository,
	authSvc AuthService,
	githubSvc GithubService,
) *User {
	return &User{
		logger: logger.With(
			slog.String("component", "service"),
			slog.String("service", "user"),
		),
		txManager: txManager,
		eventRepo: eventRepo,
		userRepo:  userRepo,
		authSvc:   authSvc,
		githubSvc: githubSvc,
	}
}

func (s *User) Me(ctx context.Context) (*model.User, error) {
	ghUser := s.authSvc.CurrentUser(ctx)

	user, err := s.userRepo.GetByGitHubID(ctx, ghUser.GetID())
	if err != nil {
		return nil, fmt.Errorf("failed to load user: %w", err)
	}

	return user, nil
}

func (s *User) Sync(ctx context.Context) (*model.User, error) {
	ghUser := s.authSvc.CurrentUser(ctx)
	entity := helper.MapGitHubUserToModel(ghUser, model.UserTypeRegular)

	if err := s.userRepo.Upsert(ctx, entity); err != nil {
		return nil, fmt.Errorf("failed to upsert user: %w", err)
	}

	go func() {
		err := s.collectAndStoreGitHubFollowers(
			context.WithoutCancel(ctx),
			entity,
		)
		if err != nil {
			s.logger.ErrorContext(
				ctx,
				"unable to store GitHub followers",
				slog.Any("user_id", entity.ID),
				slog.Any("error", err),
			)
		}
	}()

	return entity, nil
}

func (s *User) GetRegularUsers(ctx context.Context, querier query.Querier) ([]*model.User, error) {
	users, err := s.userRepo.List(
		ctx,
		repository.WithWhere("type = ?", model.UserTypeRegular),
		repository.WithOrder("id"),
		repository.WithPreload("Followers"),
		repository.WithQuerier(querier),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load users: %w", err)
	}

	return users, nil
}

func (s *User) StoreGitHubFollowers(ctx context.Context, user *model.User, followers []*github.User) error {
	logger := s.logger.With(slog.Any("user_id", user.ID))
	logger.DebugContext(ctx, "storing GitHub followers for user")

	changes := helper.NewUserChanges(
		user.Followers,
		helper.MapGitHubUsersToModels(followers, model.UserTypeReference),
	)
	logger.DebugContext(
		ctx,
		"follower changes calculated",
		slog.Bool("has_additions", changes.HasAdditions()),
		slog.Bool("has_deletions", changes.HasDeletions()),
		slog.Int("additions_count", changes.AdditionsCount()),
		slog.Int("deletions_count", changes.DeletionsCount()),
	)

	return s.processFollowerChanges(ctx, user, changes)
}

func (s *User) GetFollowers(
	ctx context.Context,
	userID uuid.UUID,
	querier query.Querier,
) ([]*model.User, error) {
	_, err := s.userRepo.GetByID(ctx, userID, repository.WithSelect("id"))
	if err != nil {
		return nil, fmt.Errorf("failed to load user: %w", err)
	}

	followers, err := s.userRepo.ListFollowers(ctx, userID, repository.WithQuerier(querier))
	if err != nil {
		return nil, fmt.Errorf("failed to load followers: %w", err)
	}

	return followers, nil
}

func (s *User) GetFollowEvents(
	ctx context.Context,
	userID uuid.UUID,
	querier query.Querier,
) ([]*model.Event, error) {
	_, err := s.userRepo.GetByID(ctx, userID, repository.WithSelect("id"))
	if err != nil {
		return nil, fmt.Errorf("failed to load user: %w", err)
	}

	events, err := s.eventRepo.List(
		ctx,
		repository.WithWhere("user_id = ?", userID),
		repository.WithPreload("ReferenceUser"),
		repository.WithQuerier(querier),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load follow events: %w", err)
	}

	return events, nil
}

func (s *User) processFollowerChanges(ctx context.Context, user *model.User, changes *helper.UserChanges) error {
	logger := s.logger.With(slog.Any("user_id", user.ID))

	if !changes.HasChanges() {
		logger.DebugContext(ctx, "no changes in user followers")

		return nil
	}

	logger.DebugContext(ctx, "starting to process follower changes")

	err := s.txManager.RunInTransaction(ctx, func(withTx repository.Option) error {
		if changes.HasAdditions() {
			if err := s.processNewFollowers(ctx, user, changes.Additions(), withTx); err != nil {
				return err
			}
		}

		if changes.HasDeletions() {
			if err := s.processRemovedFollowers(ctx, user, changes.Deletions(), withTx); err != nil {
				return err
			}
		}

		logger.DebugContext(ctx, "successfully processed all follower changes")

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to process follower changes in transaction: %w", err)
	}

	return nil
}

// If we have new followers, we have to:
// 1. Store them in the Users table;
// 2. Create a reference (association) in the UserFollowers table.
// 3. Create FOLLOW Events.
func (s *User) processNewFollowers(
	ctx context.Context,
	user *model.User,
	followers []*model.User,
	withTx repository.Option,
) error {
	logger := s.logger.With(slog.Any("user_id", user.ID))

	logger.DebugContext(ctx, "upserting new followers")

	err := s.userRepo.UpsertMany(ctx, followers, withTx)
	if err != nil {
		return fmt.Errorf("failed to store users: %w", err)
	}

	logger.DebugContext(ctx, "creating follower associations")

	user.Followers = nil

	err = s.userRepo.AddFollowers(ctx, user, followers, withTx)
	if err != nil {
		return fmt.Errorf("failed to add followers: %w", err)
	}

	events := helper.CreateUserReferenceEvents(user, followers, model.EventTypeFollow)
	if err = s.eventRepo.CreateMany(ctx, events, withTx); err != nil {
		return fmt.Errorf("failed to create follow events: %w", err)
	}

	return nil
}

// If we have unfollows, we have to remove the users from the UserFollowers table.
// We do not remove them from the Users table, as they may follow other users as well.
// Additionally, we create UNFOLLOW Events.
func (s *User) processRemovedFollowers(
	ctx context.Context,
	user *model.User,
	followers []*model.User,
	withTx repository.Option,
) error {
	logger := s.logger.With(slog.Any("user_id", user.ID))

	logger.DebugContext(ctx, "removing follower associations")

	err := s.userRepo.RemoveFollowers(ctx, user, followers, withTx)
	if err != nil {
		return fmt.Errorf("failed to remove followers: %w", err)
	}

	events := helper.CreateUserReferenceEvents(user, followers, model.EventTypeUnfollow)
	if err = s.eventRepo.CreateMany(ctx, events, withTx); err != nil {
		return fmt.Errorf("failed to create unfollow events: %w", err)
	}

	return nil
}

func (s *User) collectAndStoreGitHubFollowers(ctx context.Context, user *model.User) error {
	const ghLimit = 100

	logger := s.logger.With(slog.Any("user_id", user.ID))

	logger.InfoContext(ctx, "storing user followers")

	followers, err := s.githubSvc.
		WithToken(s.authSvc.Token(ctx)).
		CollectUserFollowers(ctx, user.Username, ghLimit)
	if err != nil {
		return fmt.Errorf("failed to collect user followers: %w", err)
	}

	err = s.StoreGitHubFollowers(ctx, user, followers)
	if err != nil {
		return fmt.Errorf("failed to store GitHub followers: %w", err)
	}

	logger.DebugContext(ctx, "user followers storing process has been finished")

	return nil
}
