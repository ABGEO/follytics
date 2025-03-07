package response

import (
	"time"

	"github.com/google/uuid"
)

type EventWithUserReference struct {
	ID            uuid.UUID              `binding:"required" example:"01955908-d43b-7900-8f5c-5faa67dab4d3" json:"id"`
	CreatedAt     time.Time              `binding:"required" example:"1970-01-01T00:00:00.000+04:00"        json:"createdAt"` //nolint: lll
	Type          string                 `binding:"required" example:"FOLLOW"                               json:"type"`
	ReferenceUser *UserForEventReference `json:"user"`
}
