package model

import "github.com/google/uuid"

type EventType string

const (
	EventTypeFollow   EventType = "FOLLOW"
	EventTypeUnfollow EventType = "UNFOLLOW"
)

type Event struct {
	Base

	Type            EventType `gorm:"type:event_type"`
	UserID          uuid.UUID `gorm:"index"`
	User            *User
	ReferenceUserID uuid.UUID `gorm:"index"`
	ReferenceUser   *User
}
