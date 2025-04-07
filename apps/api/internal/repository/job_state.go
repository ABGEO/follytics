package repository

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/abgeo/follytics/internal/model"
)

type JobStateRepository interface {
	Get(ctx context.Context, jobName string, opts ...Option) (*model.JobState, error)
	Upsert(ctx context.Context, entity *model.JobState, opts ...Option) error
}

type JobState struct {
	db *gorm.DB
}

var _ JobStateRepository = (*JobState)(nil)

func NewJobState(db *gorm.DB) *JobState {
	return &JobState{
		db: db,
	}
}

func (r *JobState) Get(ctx context.Context, jobName string, opts ...Option) (*model.JobState, error) {
	var jobState *model.JobState

	tx := WithOptions(r.db, opts...)

	return jobState, tx.
		WithContext(ctx).
		Where("job_name = ?", jobName).
		First(&jobState).
		Error
}

func (r *JobState) Upsert(ctx context.Context, entity *model.JobState, opts ...Option) error {
	tx := WithOptions(r.db, opts...)

	return tx.
		WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "job_name"}},
			UpdateAll: true,
		}).Create(&entity).Error
}
