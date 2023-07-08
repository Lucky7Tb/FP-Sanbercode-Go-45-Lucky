package models

import "time"

type Follower struct {
	ID           uint      `json:"id,omitempty" gorm:"primary_key"`
	UserId       uint      `json:"user_id,omitempty"`
	FollowUserId uint      `json:"follow_user_id,omitempty"`
	User         User      `json:"user,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
