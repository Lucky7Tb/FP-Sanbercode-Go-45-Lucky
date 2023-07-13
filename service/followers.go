package service

import (
	"tulisaja/models"

	"gorm.io/gorm"
)

func GetFollowers(db *gorm.DB, userId int) ([]models.Follower, error) {
	var followersUser []models.Follower
	if err := db.Select("id", "user_id").Where("follow_user_id = ?", userId).Find(&followersUser).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(followersUser); i++ {
		var user models.User
		db.Select("username").Where("id = ?", followersUser[i].UserId).First(&user)
		followersUser[i].User = &user
	}
	return followersUser, nil
}
