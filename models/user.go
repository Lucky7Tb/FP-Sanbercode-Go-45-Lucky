package models

type User struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	FullName  string     `json:"full_name,omitempty"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"-"`
	Article   []Article  `json:"article,omitempty"`
	Followers []Follower `json:"followers,omitempty"`
}
