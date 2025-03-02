package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/abgeo/follytics/internal/domain/constant"
	"github.com/abgeo/follytics/internal/domain/dto/response"
)

type HTTPService interface {
	HTTPResponse(ctx *gin.Context, httpCode int, data any)
	HTTPError(ctx *gin.Context, httpCode int, code string, message string)
	HTTPValidationError(ctx *gin.Context, err error)
}

type HTTP struct{}

var _ HTTPService = (*HTTP)(nil)

func NewHTTP() *HTTP {
	return &HTTP{}
}

func (s *HTTP) HTTPResponse(ctx *gin.Context, httpCode int, data any) {
	ctx.JSON(httpCode, data)
}

func (s *HTTP) HTTPError(ctx *gin.Context, httpCode int, code string, message string) {
	ctx.AbortWithStatusJSON(httpCode, response.HTTPError{
		Code:    code,
		Message: message,
	})
}

func (s *HTTP) HTTPValidationError(ctx *gin.Context, err error) {
	s.HTTPError(
		ctx,
		http.StatusBadRequest,
		constant.HTTPErrorCodeInvalidPayload,
		strings.Join(s.normalizeHTTPValidationError(err), "\n"),
	)
}

func (s *HTTP) normalizeHTTPValidationError(err error) []string {
	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		return s.formatValidationErrors(validationErr)
	}

	return []string{"invalid payload"}
}

func (s *HTTP) formatValidationErrors(validationErrors validator.ValidationErrors) []string {
	messages := make([]string, 0, len(validationErrors))

	for _, fieldError := range validationErrors {
		var errorMessage string

		switch fieldError.Tag() {
		case "required":
			errorMessage = fmt.Sprintf("Key '%s' is required", fieldError.Field())
		case "email":
			errorMessage = fmt.Sprintf("Value of field '%s' is not a valid Email address", fieldError.Field())
		default:
			errorMessage = fieldError.Error()
		}

		messages = append(messages, errorMessage)
	}

	return messages
}
