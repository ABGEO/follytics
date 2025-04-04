package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log/slog"
	"maps"
	"net/http"
	"slices"
	"sort"
	"time"

	"github.com/google/uuid"

	"github.com/abgeo/follytics/internal/domain/dto"
	"github.com/abgeo/follytics/internal/model"
	"github.com/abgeo/follytics/internal/repository"
)

type EventService interface {
	FollowersTimeline(ctx context.Context, userID uuid.UUID) (*dto.FollowersTimeline, error)
}

type Event struct {
	logger    *slog.Logger
	txManager *repository.TransactionManager

	eventRepo repository.EventRepository
	userRepo  repository.UserRepository
}

var _ EventService = (*Event)(nil)

func NewEvent(
	logger *slog.Logger,
	txManager *repository.TransactionManager,
	eventRepo repository.EventRepository,
	userRepo repository.UserRepository,
) *Event {
	return &Event{
		logger: logger.With(
			slog.String("component", "service"),
			slog.String("service", "event"),
		),
		txManager: txManager,
		eventRepo: eventRepo,
		userRepo:  userRepo,
	}
}

func (s *Event) FollowersTimeline(ctx context.Context, userID uuid.UUID) (*dto.FollowersTimeline, error) {
	user, err := s.userRepo.GetByID(ctx, userID, repository.WithSelect("username,name,avatar"))
	if err != nil {
		return nil, fmt.Errorf("failed to load user: %w", err)
	}

	if user.Avatar != "" {
		// @todo: implement caching.
		user.Avatar, err = s.getUserAvatar(ctx, user.Avatar)
		if err != nil {
			user.Avatar = ""

			s.logger.With(
				slog.Any("error", err),
				slog.Any("user_id", userID),
			).Error("failed to get user avatar")
		}
	}

	startDate, endDate := s.getTimelineRange()

	// @todo: implement caching.
	events, err := s.eventRepo.AggregateEventsByDateAndType(
		ctx,
		repository.WithWhere("user_id", userID),
		repository.WithWhere("created_at BETWEEN ? AND ?", startDate, endDate),
		repository.WithOrder("date"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load aggregated events: %w", err)
	}

	return &dto.FollowersTimeline{
		User:     user,
		Timeline: s.calculateDailyFollowerChanges(events),
	}, nil
}

func (s *Event) getUserAvatar(ctx context.Context, url string) (string, error) {
	response, err := s.fetchAvatar(ctx, url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read avatar data: %w", err)
	}

	base64Data := base64.StdEncoding.EncodeToString(imageData)

	return "data:image/jpeg;base64," + base64Data, nil
}

func (s *Event) fetchAvatar(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch avatar: %w", err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch avatar: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch avatar: %w", err)
	}

	return response, nil
}

func (s *Event) getTimelineRange() (time.Time, time.Time) {
	now := time.Now()

	return now.AddDate(0, -1, 0), now
}

func (s *Event) calculateDailyFollowerChanges(events []model.AggregatedEvent) []dto.DailyFollowerChange {
	dailyChangesMap := s.calculateDailyFollowerChangesMap(events)

	result := slices.Collect(maps.Values(dailyChangesMap))
	sort.Slice(result, func(i, j int) bool {
		return result[i].Date.Before(result[j].Date)
	})

	return result
}

func (s *Event) calculateDailyFollowerChangesMap(events []model.AggregatedEvent) map[int64]dto.DailyFollowerChange {
	totalFollowers := 0
	dailyChanges := make(map[int64]dto.DailyFollowerChange, len(events))

	for _, event := range events {
		index := event.Date.Unix()

		change, ok := dailyChanges[index]
		if !ok {
			change = dto.DailyFollowerChange{
				Date: event.Date,
			}
		}

		switch event.Type {
		case model.EventTypeFollow:
			totalFollowers += event.Count
			change.Follows += event.Count
		case model.EventTypeUnfollow:
			totalFollowers -= event.Count
			change.Unfollows += event.Count
		}

		change.Total = totalFollowers
		dailyChanges[index] = change
	}

	return dailyChanges
}
