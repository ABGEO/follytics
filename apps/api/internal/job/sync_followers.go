package job

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/go-github/v68/github"
	"gorm.io/datatypes"

	"github.com/abgeo/follytics/internal/config"
	"github.com/abgeo/follytics/internal/model"
	"github.com/abgeo/follytics/internal/pagination"
	"github.com/abgeo/follytics/internal/service"
)

type SyncFollowers struct {
	logger    *slog.Logger
	conf      *config.Config
	jobConfig config.SyncFollowersJob

	githubSvc   service.GithubService
	jobStateSvc service.JobStateService
	userSvc     service.UserService
}

var _ Job = (*SyncFollowers)(nil)

func NewSyncFollowers(
	logger *slog.Logger,
	conf *config.Config,
	githubSvc service.GithubService,
	jobStateSvc service.JobStateService,
	userSvc service.UserService,
) *SyncFollowers {
	return &SyncFollowers{
		logger: logger.With(
			slog.String("component", "job"),
			slog.String("job", "sync_followers"),
		),
		conf:        conf,
		jobConfig:   conf.Worker.Job.SyncFollowers,
		githubSvc:   githubSvc,
		jobStateSvc: jobStateSvc,
		userSvc:     userSvc,
	}
}

func (j *SyncFollowers) Name() string {
	return "sync_followers"
}

func (j *SyncFollowers) Run(ctx context.Context) error {
	if err := j.initializeGitHubToken(ctx); err != nil {
		return err
	}

	offset := j.getJobOffset(ctx, 0)

	users, offset, err := j.loadUsers(ctx, offset)
	if err != nil {
		return err
	}

	j.processUsers(ctx, users)

	return j.storeJobOffset(ctx, offset+len(users))
}

func (j *SyncFollowers) initializeGitHubToken(ctx context.Context) error {
	var err error

	j.logger.DebugContext(ctx, "initializing GitHub installation token")

	j.githubSvc, err = j.githubSvc.WithInstallationToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to setup GitHub installation token: %w", err)
	}

	return nil
}

func (j *SyncFollowers) getJobOffset(ctx context.Context, defaultValue int) int {
	j.logger.DebugContext(ctx, "loading job offset from job state attributes")

	jobStateAttributes, err := j.jobStateSvc.GetAttributes(ctx, j.Name())
	if err != nil {
		j.logger.ErrorContext(ctx, "failed to load job state attributes", slog.Any("error", err))

		return defaultValue
	}

	rawValue, ok := jobStateAttributes["offset"]
	if !ok {
		j.logger.ErrorContext(ctx, "offset key not found in job state attributes")

		return defaultValue
	}

	jsonValue, ok := rawValue.(json.Number)
	if !ok {
		j.logger.ErrorContext(ctx, "offset value is not a JSON number")

		return defaultValue
	}

	value, err := jsonValue.Int64()
	if err != nil {
		j.logger.ErrorContext(ctx, "failed to convert offset to integer", slog.Any("error", err))

		return defaultValue
	}

	return int(value)
}

func (j *SyncFollowers) loadUsers(ctx context.Context, offset int) ([]*model.User, int, error) {
	j.logger.DebugContext(ctx, "loading users to process", slog.Int("offset", offset))

	paginator := pagination.New().
		WithLimit(j.jobConfig.BatchSize).
		WithOffset(offset)

	users, err := j.userSvc.GetRegularUsers(ctx, paginator)
	if err != nil {
		return nil, offset, fmt.Errorf("failed to load users: %w", err)
	}

	if len(users) == 0 && offset != 0 {
		j.logger.DebugContext(ctx, "no users to process, resetting offset")

		return j.loadUsers(ctx, 0)
	}

	return users, offset, nil
}

func (j *SyncFollowers) processUsers(ctx context.Context, users []*model.User) {
	var wg sync.WaitGroup

	j.logger.DebugContext(ctx, "starting user processing", slog.Int("count", len(users)))

	for _, user := range users {
		wg.Add(1)

		go func(user *model.User) {
			defer wg.Done()

			if err := j.processUser(ctx, user); err != nil {
				j.logger.ErrorContext(ctx, "failed to process user",
					slog.Any("user_id", user.ID),
					slog.Any("error", err),
				)
			}
		}(user)
	}

	wg.Wait()
	j.logger.DebugContext(ctx, "all users processed")
}

func (j *SyncFollowers) processUser(ctx context.Context, user *model.User) error {
	logger := j.logger.With(slog.Any("user_id", user.ID))
	logger.DebugContext(ctx, "processing user")

	followers, err := j.fetchUserFollowers(ctx, user)
	if err != nil {
		return err
	}

	err = j.userSvc.StoreGitHubFollowers(ctx, user, followers)
	if err != nil {
		return fmt.Errorf("failed to store GitHub followers: %w", err)
	}

	return nil
}

func (j *SyncFollowers) fetchUserFollowers(ctx context.Context, user *model.User) ([]*github.User, error) {
	var followers []*github.User

	page := 1
	logger := j.logger.With(slog.Any("user_id", user.ID))

	for {
		logger.DebugContext(ctx, "fetching followers", slog.Int("page", page))

		users, res, err := j.githubSvc.GetUserFollowers(ctx, user.Username, page, j.jobConfig.GitHubPageSize)
		if err != nil {
			return nil, fmt.Errorf("failed to get user followers: %w", err)
		}

		followers = append(followers, users...)

		if res.NextPage == 0 {
			logger.DebugContext(ctx, "last page reached")

			break
		}

		page = res.NextPage
	}

	logger.DebugContext(ctx, "followers fetched", slog.Int("count", len(followers)))

	return followers, nil
}

func (j *SyncFollowers) storeJobOffset(ctx context.Context, offset int) error {
	j.logger.DebugContext(ctx, "storing job offset in job state attributes", slog.Int("offset", offset))

	attributes := datatypes.JSONMap{
		"offset": offset,
	}

	if _, err := j.jobStateSvc.StoreAttributes(ctx, j.Name(), attributes); err != nil {
		return fmt.Errorf("failed to store job state attributes: %w", err)
	}

	return nil
}
