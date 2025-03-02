package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
}

func (m *Base) BeforeCreate(_ *gorm.DB) error {
	var err error

	if m.ID == uuid.Nil {
		m.ID, err = uuid.NewV7()
		if err != nil {
			return fmt.Errorf("failed to create UUID: %w", err)
		}
	}

	return nil
}
