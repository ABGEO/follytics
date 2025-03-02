package worker

import (
	"context"
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"sync"

	"github.com/abgeo/follytics/internal/domain/errors"
	"github.com/abgeo/follytics/internal/job"
)

type Worker interface {
	Process(ctx context.Context, jobNames []string) error
}

type JobWorker struct {
	logger *slog.Logger

	jobs map[string]job.Job
}

var _ Worker = (*JobWorker)(nil)

func NewJobWorker(
	logger *slog.Logger,
	syncFollowersJob job.Job,
) *JobWorker {
	return &JobWorker{
		logger: logger.With(
			slog.String("component", "worker"),
		),
		jobs: map[string]job.Job{
			syncFollowersJob.Name(): syncFollowersJob,
		},
	}
}

func (w *JobWorker) Process(ctx context.Context, jobNames []string) error {
	var wg sync.WaitGroup

	w.logger.InfoContext(ctx, "scheduling jobs", slog.Any("jobs", jobNames))

	jobs, err := w.getJobs(jobNames)
	if err != nil {
		return err
	}

	for _, jobInstance := range jobs {
		wg.Add(1)

		go func(jobInstance job.Job) {
			defer wg.Done()

			w.logger.InfoContext(
				ctx,
				"starting job",
				slog.String("job", jobInstance.Name()),
			)

			if err := jobInstance.Run(ctx); err != nil {
				w.logger.ErrorContext(
					ctx,
					"failed to execute job",
					slog.String("job", jobInstance.Name()),
					slog.Any("error", err),
				)
			}
		}(jobInstance)
	}

	w.logger.InfoContext(ctx, "jobs have been scheduled successfully, waiting for execution")
	wg.Wait()
	w.logger.InfoContext(ctx, "all jobs have been executed")

	return nil
}

func (w *JobWorker) getJobs(jobNames []string) ([]job.Job, error) {
	if slices.Contains(jobNames, "all") {
		return slices.Collect(maps.Values(w.jobs)), nil
	}

	jobs := make([]job.Job, 0, len(jobNames))
	for _, jobName := range jobNames {
		value, ok := w.jobs[jobName]
		if !ok {
			return nil, fmt.Errorf("job name '%s' is invalid: %w", jobName, errors.ErrJobIsInvalid)
		}

		jobs = append(jobs, value)
	}

	return jobs, nil
}
