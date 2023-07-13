package models

import "time"

type Follower struct {
	ID           uint      `json:"id,omitempty"`
	UserId       uint      `json:"user_id,omitempty"`
	FollowUserId uint      `json:"follow_user_id,omitempty"`
	User         *User     `json:"user,omitempty" gorm:"foreignkey:UserId;association_foreignkey:id"`
	FollowUser   *User     `json:"follow_user,omitempty" gorm:"foreignkey:FollowUserId;association_foreignkey:id"`
	CreatedAt    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP()"`
}
