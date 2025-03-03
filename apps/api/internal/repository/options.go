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

func WithPagination(offset int, limit int) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(offset).Limit(limit)
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
