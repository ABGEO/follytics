package response

import "time"

type FollowersTimelineItem struct {
	Date      time.Time `binding:"required" example:"1970-01-01T00:00:00.000+04:00" json:"date"`
	Total     int       `binding:"required" example:"1500"                          json:"total"`
	Follows   int       `binding:"required" example:"100"                           json:"follows"`
	Unfollows int       `binding:"required" example:"20"                            json:"unfollows"`
}

type FollowersTimeline struct {
	User     UserForTimeline         `binding:"required" json:"user"`
	Timeline []FollowersTimelineItem `binding:"required" json:"timeline"`
}
