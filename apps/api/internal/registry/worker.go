package registry

import (
	"github.com/spf13/pflag"

	"github.com/abgeo/follytics/internal/job"
	"github.com/abgeo/follytics/internal/repository"
	"github.com/abgeo/follytics/internal/service"
	workerpkg "github.com/abgeo/follytics/internal/worker"
)

type WorkerRegistry interface {
	Registry

	GetWorker() *workerpkg.JobWorker
}

type Worker struct {
	Registry

	worker *workerpkg.JobWorker

	syncFollowersJob job.Job

	eventRepo    repository.EventRepository
	jobStateRepo repository.JobStateRepository
	userRepo     repository.UserRepository

	authSvc     service.AuthService
	githubSvc   service.GithubService
	jobStateSvc service.JobStateService
	userSvc     service.UserService
}

var _ WorkerRegistry = (*Worker)(nil)

func NewWorker(flags *pflag.FlagSet) (*Worker, error) {
	reg := &Worker{}

	baseRegistry, err := NewBase(flags)
	if err != nil {
		return nil, err
	}

	reg.Registry = baseRegistry
	reg.createDependencies()

	return reg, nil
}

func (r *Worker) GetWorker() *workerpkg.JobWorker {
	return r.worker
}

func (r *Worker) createDependencies() {
	r.eventRepo = repository.NewEvent(r.GetDB())
	r.jobStateRepo = repository.NewJobState(r.GetDB())
	r.userRepo = repository.NewUser(r.GetDB())

	r.authSvc = service.NewAuth()
	r.githubSvc = service.NewGithub(r.GetConfig(), r.GetLogger())
	r.jobStateSvc = service.NewJobState(r.GetLogger(), r.jobStateRepo)
	r.userSvc = service.NewUser(r.GetLogger(), r.GetTransactionManager(), r.eventRepo, r.userRepo, r.authSvc, r.githubSvc)

	r.syncFollowersJob = job.NewSyncFollowers(
		r.GetLogger(),
		r.GetConfig(),
		r.githubSvc,
		r.jobStateSvc,
		r.userSvc,
	)

	r.worker = workerpkg.NewJobWorker(r.GetLogger(), r.syncFollowersJob)
}
