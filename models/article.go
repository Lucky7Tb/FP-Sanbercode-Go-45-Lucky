package models

type Article struct {
	ID        uint      `json:"id,omitempty" gorm:"primary_key"`
	UserId    uint      `json:"user_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	User      *User     `json:"user,omitempty"`
	Comment   []Comment `json:"comment,omitempty"`
	Likes     []Like    `json:"like,omitempty"`
	CreatedAt string    `json:"created_at,omitempty"`
	UpdatedAt string    `json:"updated_at,omitempty"`
}
