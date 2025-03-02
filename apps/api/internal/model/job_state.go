package model

import "gorm.io/datatypes"

type JobState struct {
	JobName    string `gorm:"primaryKey"`
	Attributes datatypes.JSONMap
}
