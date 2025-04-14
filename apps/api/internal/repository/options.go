package repository

import (
	"gorm.io/gorm"

	"github.com/abgeo/follytics/internal/query"
	"github.com/abgeo/follytics/internal/query/filter"
	"github.com/abgeo/follytics/internal/query/pagination"
)

type Option func(tx *gorm.DB) *gorm.DB

func WithOptions(tx *gorm.DB, opts ...Option) *gorm.DB {
	for _, opt := range opts {
		tx = opt(tx)
	}

	return tx
}

func WithQuerier(querier query.Querier) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Scopes(querier.Apply)
	}
}

func WithPagination(paginator pagination.Paginator) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Scopes(paginator.Apply)
	}
}

func WithFilterer(filterer filter.Filterer) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Scopes(filterer.Apply)
	}
}

func WithSelect(query interface{}, args ...interface{}) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Select(query, args...)
	}
}

func WithJoins(query string, args ...interface{}) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Joins(query, args...)
	}
}

func WithWhere(query interface{}, args ...interface{}) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query, args...)
	}
}

func WithOrder(value interface{}) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(value)
	}
}

func WithPreload(query string, args ...interface{}) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Preload(query, args...)
	}
}

func WithTransaction(tx *gorm.DB) func(tx *gorm.DB) *gorm.DB {
	return func(_ *gorm.DB) *gorm.DB {
		return tx
	}
}

func WithDebug() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Debug()
	}
}
