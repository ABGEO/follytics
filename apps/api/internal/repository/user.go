package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/abgeo/follytics/internal/model"
)

type UserRepository interface {
	List(ctx context.Context, opts ...Option) ([]*model.User, error)
	Get(ctx context.Context, opts ...Option) (*model.User, error)
	GetByID(ctx context.Context, id uuid.UUID, opts ...Option) (*model.User, error)
	GetByGitHubID(ctx context.Context, gitHubID int64, opts ...Option) (*model.User, error)
	Upsert(ctx context.Context, entity *model.User, opts ...Option) error
	UpsertMany(ctx context.Context, entities []*model.User, opts ...Option) error
	AddFollowers(ctx context.Context, entity *model.User, followers []*model.User, opts ...Option) error
	RemoveFollowers(ctx context.Context, entity *model.User, followers []*model.User, opts ...Option) error
	ListFollowers(ctx context.Context, userID uuid.UUID, offset int, limit int, opts ...Option) ([]*model.User, error)
}

type User struct {
	db *gorm.DB
}

var _ UserRepository = (*User)(nil)

func NewUser(db *gorm.DB) *User {
	// @todo: move to separate migrator tool.
	db.Exec(`
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_type') THEN
		CREATE TYPE user_type AS ENUM('REGULAR', 'REFERENCE');
    END IF;
END $$;
`)
	db.AutoMigrate(&model.User{}) //nolint: errcheck

	return &User{
		db: db,
	}
}

func (r *User) List(ctx context.Context, opts ...Option) ([]*model.User, error) {
	var users []*model.User

	tx := WithOptions(r.db, opts...)

	return users, tx.
		WithContext(ctx).
		Find(&users).
		Error
}

func (r *User) Get(ctx context.Context, opts ...Option) (*model.User, error) {
	var user *model.User

	tx := WithOptions(r.db, opts...)

	return user, tx.
		WithContext(ctx).
		First(&user).
		Error
}

func (r *User) GetByID(ctx context.Context, id uuid.UUID, opts ...Option) (*model.User, error) {
	opts = append(opts, WithWhere(id))

	return r.Get(ctx, opts...)
}

func (r *User) GetByGitHubID(ctx context.Context, gitHubID int64, opts ...Option) (*model.User, error) {
	opts = append(opts, WithWhere("gh_id = ?", gitHubID))

	return r.Get(ctx, opts...)
}

func (r *User) Upsert(ctx context.Context, entity *model.User, opts ...Option) error {
	tx := WithOptions(r.db, opts...)

	return tx.WithContext(ctx).
		Clauses(
			clause.Returning{
				Columns: []clause.Column{{Name: "id"}},
			},
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "gh_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"username", "name", "email", "avatar", "type"}),
			},
		).
		Create(&entity).
		Error
}

func (r *User) UpsertMany(ctx context.Context, entities []*model.User, opts ...Option) error {
	tx := WithOptions(r.db, opts...)

	return tx.WithContext(ctx).
		Clauses(
			clause.Returning{
				Columns: []clause.Column{{Name: "id"}},
			},
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "gh_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"username", "name", "email", "avatar"}),
			},
		).
		Create(&entities).
		Error
}

func (r *User) AddFollowers(ctx context.Context, entity *model.User, followers []*model.User, opts ...Option) error {
	tx := WithOptions(r.db, opts...)

	err := tx.Model(&entity).
		WithContext(ctx).
		Omit("Followers.*").
		Association("Followers").
		Append(followers)
	if err != nil {
		return fmt.Errorf("failed to append assiciation: %w", err)
	}

	return nil
}

func (r *User) RemoveFollowers(
	ctx context.Context,
	entity *model.User,
	followers []*model.User,
	opts ...Option,
) error {
	tx := WithOptions(r.db, opts...)

	err := tx.Model(&entity).
		WithContext(ctx).
		Association("Followers").
		Delete(followers)
	if err != nil {
		return fmt.Errorf("failed to delete assiciation: %w", err)
	}

	return nil
}

func (r *User) ListFollowers(
	ctx context.Context,
	userID uuid.UUID,
	offset int,
	limit int,
	opts ...Option,
) ([]*model.User, error) {
	opts = append(
		opts,
		WithSelect(`"user".*`),
		WithJoins(`JOIN user_followers uf on "user".id = uf.follower_id`),
		WithWhere("uf.user_id = ?", userID),
		WithOrder("id"),
		WithPagination(offset, limit),
	)

	return r.List(ctx, opts...)
}
