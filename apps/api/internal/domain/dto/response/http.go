package response

import (
	"fmt"

	"github.com/jinzhu/copier"

	domainErrors "github.com/abgeo/follytics/internal/domain/errors"
)

type HTTPResponse[T any] struct {
	Data T `binding:"required" json:"data"`
}

func (resp *HTTPResponse[T]) Populate(raw any) error {
	if raw == nil {
		return domainErrors.ErrCannotPopulateNil
	}

	if err := copier.Copy(&resp.Data, raw); err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}

	return nil
}

type HTTPError struct {
	Code    string `binding:"required" example:"UNKNOWN"            json:"code"`
	Message string `binding:"required" example:"some unknown error" json:"message"`
}
