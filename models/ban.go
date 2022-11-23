package models

import (
	"time"

	"gorm.io/gorm"
)

type Ban struct {
	gorm.Model
	UserID    int       `json:"userId" form:"userId"`
	TopicID   int       `json:"topicId" form:"topicId"`
	Ban_Until time.Time `json:"ban_until" form:"ban_until"`

	User  User
	Topic Topic
}

func (Ban) TableName() string {
	return "bans"
}
