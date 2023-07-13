package service

import (
	"tulisaja/models"

	"gorm.io/gorm"
)

func GetFollowingUser(db *gorm.DB, userId int) ([]models.Follower, error) {
	var followingUser []models.Follower
	if err := db.Select("id", "follow_user_id").Where("user_id = ?", userId).Find(&followingUser).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(followingUser); i++ {
		var user models.User
		db.Select("username").Where("id = ?", followingUser[i].FollowUserId).First(&user)
		followingUser[i].FollowUser = &user
	}
	return followingUser, nil
}

func FollowUser(db *gorm.DB, userId int, username string) error {
	var user models.User

	if err := db.Where("username = ?", username).Take(&user).Error; err != nil {
		return err
	}

	return db.Create(&models.Follower{UserId: uint(userId), FollowUserId: user.ID}).Error
}

func DeleteFollowingUser(db *gorm.DB, userId, followingId int) error {
	var followers models.Follower
	if err := db.Where("id = ? AND user_id = ?", followingId, userId).Take(&followers).Error; err != nil {
		return err
	}
	db.Delete(&followers)
	return nil
}
