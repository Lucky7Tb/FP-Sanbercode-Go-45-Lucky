package models

type Article struct {
	ID      uint      `json:"id,omitempty" gorm:"primary_key"`
	UserId  uint      `json:"user_id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	User    *User     `json:"user,omitempty"`
	Comment []Comment `json:"comment,omitempty"`
	Likes   int       `json:"like" gorm:"-"`
}
