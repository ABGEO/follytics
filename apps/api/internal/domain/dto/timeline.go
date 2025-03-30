package dto

import (
	"time"

	"github.com/abgeo/follytics/internal/model"
)

type DailyFollowerChange struct {
	Date      time.Time
	Total     int
	Follows   int
	Unfollows int
}

type FollowersTimeline struct {
	User     *model.User
	Timeline []DailyFollowerChange
}
