package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type TransactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{
		db: db,
	}
}

func (tm *TransactionManager) RunInTransaction(ctx context.Context, fn func(withTx Option) error) error {
	err := tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if txErr := fn(WithTransaction(tx)); txErr != nil {
			return fmt.Errorf("transaction function failed: %w", txErr)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("database transaction failed: %w", err)
	}

	return nil
}
