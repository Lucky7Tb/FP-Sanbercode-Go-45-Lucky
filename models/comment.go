package models

import "time"

type Comment struct {
	ID        uint      `json:"id,omitempty" gorm:"primary_key"`
	UserId    uint      `json:"user_id,omitempty"`
	ArticleId uint      `json:"article_id,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	User      User      `json:"user,omitempty"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP()"`
}
