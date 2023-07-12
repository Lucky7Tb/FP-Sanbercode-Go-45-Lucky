package models

type Follower struct {
	ID           uint   `json:"id,omitempty" gorm:"primary_key"`
	UserId       uint   `json:"user_id,omitempty"`
	FollowUserId uint   `json:"follow_user_id,omitempty"`
	User         User   `json:"user,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}
