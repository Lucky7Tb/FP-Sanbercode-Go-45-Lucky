package models

import "time"

type User struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	FullName  string     `json:"full_name,omitempty"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"-"`
	Article   []Article  `json:"article,omitempty"`
	Followers []Follower `json:"followers,omitempty"`
	CreatedAt time.Time  `json:"-" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time  `json:"-" gorm:"default:CURRENT_TIMESTAMP()"`
}
