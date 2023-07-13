package models

import "time"

type Like struct {
	ID        uint      `json:"id,omitempty" gorm:"primary_key"`
	ArticleId uint      `json:"article_id,omitempty"`
	Counter   uint      `json:"counter,omitempty"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP()"`
}
