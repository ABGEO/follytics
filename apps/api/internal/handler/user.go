package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/abgeo/follytics/internal/domain/constant"
	"github.com/abgeo/follytics/internal/domain/dto/response"
	"github.com/abgeo/follytics/internal/service"
)

type UserHandler interface {
	Me(ctx *gin.Context)
	TrackLogin(ctx *gin.Context)
}

type User struct {
	logger *slog.Logger

	authSvc service.AuthService
	httpSvc service.HTTPService
	userSvc service.UserService
}

var _ UserHandler = (*User)(nil)

func NewUser(
	logger *slog.Logger,
	authSvc service.AuthService,
	httpSvc service.HTTPService,
	userSvc service.UserService,
) *User {
	return &User{
		logger: logger.With(
			slog.String("component", "handler"),
			slog.String("handler", "user"),
		),
		authSvc: authSvc,
		httpSvc: httpSvc,
		userSvc: userSvc,
	}
}

// Me handler.
//
//	@Summary		Get current account
//	@Description	Get current authenticated account
//	@ID				getCurrentUser
//	@Tags			User
//	@Security		ApiKeyAuth
//	@Success		200	{object}	response.HTTPResponse[response.User]
//	@Failure		401	{object}	response.HTTPError
//	@Failure		500	{object}	response.HTTPError
//	@Router			/users/me [get]
func (h *User) Me(ctx *gin.Context) {
	var resp response.HTTPResponse[response.User]

	user, err := h.userSvc.Me(ctx)
	if err != nil {
		h.httpSvc.HTTPError(ctx, http.StatusInternalServerError, constant.HTTPErrorCodeUnknown, err.Error())

		return
	}

	if err = resp.Populate(user); err != nil {
		h.httpSvc.HTTPError(ctx, http.StatusInternalServerError, constant.HTTPErrorCodeUnknown, err.Error())

		return
	}

	h.httpSvc.HTTPResponse(ctx, http.StatusOK, resp)
}

// TrackLogin handler.
//
//	@Summary		Track user login
//	@Description	Update the user or create new one if it does not exist
//	@ID				trackLogin
//	@Tags			User
//	@Security		ApiKeyAuth
//	@Success		200	{object}	response.HTTPResponse[response.User]
//	@Failure		400	{object}	response.HTTPError
//	@Failure		401	{object}	response.HTTPError
//	@Failure		500	{object}	response.HTTPError
//	@Router			/users/login-events [post]
func (h *User) TrackLogin(ctx *gin.Context) {
	var resp response.HTTPResponse[response.User]

	user, err := h.userSvc.Sync(ctx)
	if err != nil {
		h.httpSvc.HTTPError(ctx, http.StatusInternalServerError, constant.HTTPErrorCodeUnknown, err.Error())

		return
	}

	if err = resp.Populate(user); err != nil {
		h.httpSvc.HTTPError(ctx, http.StatusInternalServerError, constant.HTTPErrorCodeUnknown, err.Error())

		return
	}

	h.httpSvc.HTTPResponse(ctx, http.StatusOK, resp)
}
