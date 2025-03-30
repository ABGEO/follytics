package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/abgeo/follytics/internal/model"
)

type EventRepository interface {
	CreateMany(ctx context.Context, entities []*model.Event, opts ...Option) error
	List(ctx context.Context, opts ...Option) ([]*model.Event, error)
	AggregateEventsByDateAndType(ctx context.Context, opts ...Option) ([]model.AggregatedEvent, error)
}

type Event struct {
	db *gorm.DB
}

var _ EventRepository = (*Event)(nil)

func NewEvent(db *gorm.DB) *Event {
	// @todo: move to separate migrator tool.
	db.Exec(`
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'event_type') THEN
		CREATE TYPE event_type AS ENUM('FOLLOW', 'UNFOLLOW');
    END IF;
END $$;
`)
	db.AutoMigrate(&model.Event{}) //nolint: errcheck

	return &Event{
		db: db,
	}
}

func (r *Event) CreateMany(ctx context.Context, entities []*model.Event, opts ...Option) error {
	tx := WithOptions(r.db, opts...)

	return tx.
		WithContext(ctx).
		Create(&entities).
		Error
}

func (r *Event) List(ctx context.Context, opts ...Option) ([]*model.Event, error) {
	var events []*model.Event

	tx := r.db.WithContext(ctx).
		Model(&events)

	return events, WithOptions(tx, opts...).
		Find(&events).
		Error
}

func (r *Event) AggregateEventsByDateAndType(ctx context.Context, opts ...Option) ([]model.AggregatedEvent, error) {
	var events []model.AggregatedEvent

	tx := r.db.WithContext(ctx).
		Model(&model.Event{})

	return events, WithOptions(tx, opts...).
		Select("DATE(created_at) AS date, COUNT(id) AS count, type").
		Group("DATE(created_at)").
		Group("type").
		Scan(&events).
		Error
}
