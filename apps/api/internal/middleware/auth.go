package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/abgeo/follytics/internal/domain/constant"
	"github.com/abgeo/follytics/internal/service"
)

type Auth struct {
	logger *slog.Logger

	githubSvc service.GithubService
	httpSvc   service.HTTPService
}

var _ Handler = (*Auth)(nil)

func NewAuth(logger *slog.Logger, githubSvc service.GithubService, httpSvc service.HTTPService) *Auth {
	return &Auth{
		logger: logger.With(
			slog.String("component", "middleware"),
			slog.String("middleware", "auth"),
		),
		githubSvc: githubSvc,
		httpSvc:   httpSvc,
	}
}

func (m *Auth) Handle(ctx *gin.Context) {
	accessDeniedResponse := func() {
		m.httpSvc.HTTPError(
			ctx,
			http.StatusUnauthorized,
			constant.HTTPErrorCodeUnauthorized,
			"access denied",
		)
	}

	token := ctx.GetHeader(constant.AuthTokenHeader)
	if token == "" {
		accessDeniedResponse()

		return
	}

	user, _, err := m.githubSvc.WithToken(token).GetUser(ctx, "")
	if err != nil {
		accessDeniedResponse()

		return
	}

	ctx.Set(constant.AuthTokenKey, token)
	ctx.Set(constant.AuthUserKey, user)
	ctx.Next()
}
