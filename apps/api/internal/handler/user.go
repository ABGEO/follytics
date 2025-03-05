package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/abgeo/follytics/internal/domain/constant"
	"github.com/abgeo/follytics/internal/domain/dto/response"
	"github.com/abgeo/follytics/internal/pagination"
	"github.com/abgeo/follytics/internal/service"
)

type UserHandler interface {
	Me(ctx *gin.Context)
	TrackLogin(ctx *gin.Context)
	Followers(ctx *gin.Context)
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
//	@Summary		Retrieve the current authenticated user
//	@Description	Returns details of the currently authenticated account
//	@ID				getCurrentUser
//	@Tags			User
//
//	@Security		ApiKeyAuth
//
//	@Success		200	{object}	response.HTTPResponse[response.User]
//	@Failure		401	{object}	response.HTTPError
//	@Failure		500	{object}	response.HTTPError
//
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
//	@Description	Creates a new user if they do not exist or updates an existing user upon login
//	@ID				trackLogin
//	@Tags			User
//
//	@Security		ApiKeyAuth
//
//	@Success		200	{object}	response.HTTPResponse[response.User]
//	@Failure		401	{object}	response.HTTPError
//	@Failure		500	{object}	response.HTTPError
//
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

// Followers handler.
//
//	@Summary		Retrieve followers for a user
//	@Description	Returns a paginated list of followers for the specified user ID
//	@ID				getUserFollowers
//	@Tags			User
//
//	@Security		ApiKeyAuth
//
//	@Param			page	query		string	false	"Page number for pagination (default: 1)"	Format(int)
//	@Param			limit	query		string	false	"Number of results per page (default: 10)"	Format(int)
//	@Param			id		path		string	true	"User ID to retrieve followers for"			Format(uuid)
//
//	@Success		200		{object}	response.HTTPResponse[[]response.User]{pagination=pagination.Metadata}
//	@Failure		400		{object}	response.HTTPError
//	@Failure		401		{object}	response.HTTPError
//	@Failure		404		{object}	response.HTTPError
//	@Failure		500		{object}	response.HTTPError
//
//	@Router			/users/{id}/followers [get]
func (h *User) Followers(ctx *gin.Context) {
	var resp response.HTTPResponse[[]response.User]

	paginator := pagination.New().FromContext(ctx)

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		h.httpSvc.HTTPValidationError(ctx, err)

		return
	}

	user, err := h.userSvc.GetFollowers(ctx, id, paginator)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorCode := constant.HTTPErrorCodeUnknown

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			errorCode = constant.HTTPErrorCodeNotFound
		}

		h.httpSvc.HTTPError(ctx, statusCode, errorCode, err.Error())

		return
	}

	if err = resp.PopulateWithPagination(user, paginator); err != nil {
		h.httpSvc.HTTPError(ctx, http.StatusInternalServerError, constant.HTTPErrorCodeUnknown, err.Error())

		return
	}

	h.httpSvc.HTTPResponse(ctx, http.StatusOK, resp)
}
