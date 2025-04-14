package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/abgeo/follytics/internal/model"
	"github.com/abgeo/follytics/internal/repository"
)

type JobStateService interface {
	Get(ctx context.Context, jobName string) (*model.JobState, error)
	GetOrInit(ctx context.Context, jobName string) (*model.JobState, error)
	GetAttributes(ctx context.Context, jobName string) (datatypes.JSONMap, error)
	StoreAttributes(ctx context.Context, jobName string, attributes datatypes.JSONMap) (*model.JobState, error)
}

type JobState struct {
	logger *slog.Logger

	jobStateRepo repository.JobStateRepository
}

var _ JobStateService = (*JobState)(nil)

func NewJobState(
	logger *slog.Logger,
	jobStateRepo repository.JobStateRepository,
) *JobState {
	return &JobState{
		logger: logger.With(
			slog.String("component", "service"),
			slog.String("service", "job_state"),
		),
		jobStateRepo: jobStateRepo,
	}
}

func (s *JobState) Get(ctx context.Context, jobName string) (*model.JobState, error) {
	s.logger.DebugContext(ctx, "loading job state", slog.String("job", jobName))

	jobState, err := s.jobStateRepo.Get(ctx, jobName)
	if err != nil {
		return nil, fmt.Errorf("failed to load job state: %w", err)
	}

	return jobState, nil
}

func (s *JobState) GetOrInit(ctx context.Context, jobName string) (*model.JobState, error) {
	jobState, err := s.Get(ctx, jobName)
	if err == nil {
		return jobState, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.DebugContext(ctx, "failed to find job state, initializing new one")

		return s.getEmptyJobState(jobName), nil
	}

	return nil, err
}

func (s *JobState) GetAttributes(ctx context.Context, jobName string) (datatypes.JSONMap, error) {
	jobState, err := s.GetOrInit(ctx, jobName)
	if err != nil {
		return nil, err
	}

	return jobState.Attributes, nil
}

func (s *JobState) StoreAttributes(
	ctx context.Context,
	jobName string,
	attributes datatypes.JSONMap,
) (*model.JobState, error) {
	var err error

	jobState := s.getEmptyJobState(jobName)
	jobState.Attributes = attributes

	if err = s.jobStateRepo.Upsert(ctx, jobState); err != nil {
		return nil, fmt.Errorf("failed to store job state: %w", err)
	}

	return jobState, nil
}

func (s *JobState) getEmptyJobState(jobName string) *model.JobState {
	return &model.JobState{
		JobName:    jobName,
		Attributes: make(datatypes.JSONMap),
	}
}
