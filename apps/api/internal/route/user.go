package route

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/abgeo/follytics/internal/handler"
	"github.com/abgeo/follytics/internal/middleware"
)

type User struct {
	logger         *slog.Logger
	handler        handler.UserHandler
	authMiddleware middleware.Handler
}

var _ Registerer = (*User)(nil)

func NewUser(
	logger *slog.Logger,
	handler handler.UserHandler,
	authMiddleware middleware.Handler,
) *User {
	return &User{
		logger: logger.With(
			slog.String("component", "route"),
			slog.String("route", "user"),
		),
		handler:        handler,
		authMiddleware: authMiddleware,
	}
}

func (route *User) Register(router gin.IRouter) {
	//nolint:noctx
	route.logger.Debug("setting up route")

	group := router.Group("/users")

	group.Use(route.authMiddleware.Handle)
	{
		group.GET("me", route.handler.Me)
		group.POST("login-events", route.handler.TrackLogin)
		group.GET(":id/followers", route.handler.Followers)
		group.GET(":id/follow-events", route.handler.FollowEvents)
	}

	router.GET("/users/:id/followers/timeline", route.handler.Timeline)
}
