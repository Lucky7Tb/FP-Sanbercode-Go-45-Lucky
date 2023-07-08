package models

import "time"

type Article struct {
	ID        uint      `json:"id,omitempty" gorm:"primary_key"`
	UserId    uint      `json:"user_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	Comment   []Comment `json:"comment,omitempty"`
	Likes     []Like    `json:"like,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
