package response

import (
	"fmt"

	"github.com/jinzhu/copier"

	domainErrors "github.com/abgeo/follytics/internal/domain/errors"
	"github.com/abgeo/follytics/internal/pagination"
)

type HTTPResponse[T any] struct {
	Data       T                    `binding:"required"          json:"data"`
	Pagination *pagination.Metadata `json:"pagination,omitempty" swaggerignore:"true"`
}

func (r *HTTPResponse[T]) Populate(raw any) error {
	if raw == nil {
		return domainErrors.ErrCannotPopulateNil
	}

	if err := copier.Copy(&r.Data, raw); err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}

	r.Pagination = nil

	return nil
}

func (r *HTTPResponse[T]) PopulateWithPagination(raw any, paginator pagination.Paginator) error {
	if err := r.Populate(raw); err != nil {
		return err
	}

	if paginator != nil {
		r.Pagination = paginator.GetMetadata()
	}

	return nil
}

type HTTPError struct {
	Code    string `binding:"required" example:"UNKNOWN"            json:"code"`
	Message string `binding:"required" example:"some unknown error" json:"message"`
}
