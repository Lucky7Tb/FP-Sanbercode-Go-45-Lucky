package models

import "time"

type Comment struct {
	ID        uint      `json:"id,omitempty" gorm:"primary_key"`
	UserId    uint      `json:"user_id,omitempty"`
	ArticleId uint      `json:"article_id,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
