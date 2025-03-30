package registry

import (
	"fmt"
	"net/http"

	"github.com/spf13/pflag"

	"github.com/abgeo/follytics/internal/handler"
	"github.com/abgeo/follytics/internal/middleware"
	"github.com/abgeo/follytics/internal/repository"
	"github.com/abgeo/follytics/internal/route"
	"github.com/abgeo/follytics/internal/server"
	"github.com/abgeo/follytics/internal/service"
)

type ServeRegistry interface {
	Registry

	GetRestServer() *http.Server
}

type Serve struct {
	Registry

	restServer *http.Server

	eventRepo repository.EventRepository
	userRepo  repository.UserRepository

	authSvc   service.AuthService
	eventSvc  service.EventService
	githubSvc service.GithubService
	httpSvc   service.HTTPService
	userSvc   service.UserService

	userHandler handler.UserHandler

	authMiddleware middleware.Handler

	routes []route.Registerer
}

var _ ServeRegistry = (*Serve)(nil)

func NewServe(flags *pflag.FlagSet) (*Serve, error) {
	reg := &Serve{}

	baseRegistry, err := newBase(flags)
	if err != nil {
		return nil, err
	}

	reg.Registry = baseRegistry
	reg.createDependencies()
	reg.createRoutes()

	reg.restServer, err = server.NewRest(reg.GetLogger(), reg.GetConfig(), reg.GetDB(), reg.routes)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize REST Server: %w", err)
	}

	return reg, nil
}

func (r *Serve) GetRestServer() *http.Server {
	return r.restServer
}

func (r *Serve) createDependencies() {
	r.eventRepo = repository.NewEvent(r.GetDB())
	r.userRepo = repository.NewUser(r.GetDB())

	r.authSvc = service.NewAuth()
	r.eventSvc = service.NewEvent(r.GetLogger(), r.GetTransactionManager(), r.eventRepo, r.userRepo)
	r.githubSvc = service.NewGithub(r.GetConfig())
	r.httpSvc = service.NewHTTP()
	r.userSvc = service.NewUser(r.GetLogger(), r.GetTransactionManager(), r.eventRepo, r.userRepo, r.authSvc)

	r.userHandler = handler.NewUser(r.GetLogger(), r.authSvc, r.eventSvc, r.httpSvc, r.userSvc)

	r.authMiddleware = middleware.NewAuth(r.GetLogger(), r.githubSvc, r.httpSvc)
}

func (r *Serve) createRoutes() {
	r.routes = []route.Registerer{
		route.NewUser(r.GetLogger(), r.userHandler, r.authMiddleware),
	}
}
