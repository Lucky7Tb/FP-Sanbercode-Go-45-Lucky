package models

type Comment struct {
	ID        uint   `json:"id,omitempty" gorm:"primary_key"`
	UserId    uint   `json:"user_id,omitempty"`
	ArticleId uint   `json:"article_id,omitempty"`
	Comment   string `json:"comment,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
