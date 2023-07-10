package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"tulisaja/models"
	requestinput "tulisaja/request-input/profile"
)

func ChangePassword(db *gorm.DB, input requestinput.ChangePasswordInput, userId interface{}) error {
	var user models.User
	db.Where("id = ?", userId).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		return errors.New("old password error")
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		return errors.New("failed hash password")
	}

	db.Model(&user).Where("id = ?", userId).Update("password", string(newPasswordHash))

	return nil
}
