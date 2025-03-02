package repository

import (
	"gorm.io/gorm"
)

type Option func(tx *gorm.DB) *gorm.DB

func WithOptions(tx *gorm.DB, opts ...Option) *gorm.DB {
	for _, opt := range opts {
		tx = opt(tx)
	}

	return tx
}

func WithPagination(offset int, limit int) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(offset).Limit(limit)
	}
}

func WithWhere(query interface{}, args ...interface{}) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query, args...)
	}
}

func WithOrder(value interface{}) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(value)
	}
}

func WithPreload(query string, args ...interface{}) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Preload(query, args...)
	}
}

func WithTransaction(tx *gorm.DB) Option {
	return func(_ *gorm.DB) *gorm.DB {
		return tx
	}
}

func WithDebug() Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Debug()
	}
}
