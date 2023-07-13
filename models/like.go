package models

type Like struct {
	ID        uint `json:"id,omitempty" gorm:"primary_key"`
	ArticleId uint `json:"article_id,omitempty"`
	Counter   uint `json:"counter,omitempty"`
}
