package service

import (
	"errors"
	"fmt"
	"os"
	"tulisaja/models"
	requestinput "tulisaja/request-input/auth"

	jwt "github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func Login(db *gorm.DB, input requestinput.LoginInput) (string, error) {
	var user models.User

	err := db.Where("username = ?", input.Username).First(&user).Error

	if err != nil {
		if err.Error() == "record not found" {
			return "", errors.New("not found")
		}
		return "", errors.New(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"username":  user.Username,
		"full_name": user.FullName,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New(err.Error())
	}
	return tokenString, nil
}

func Register(db *gorm.DB, input requestinput.RegisterInput) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return errors.New("Failed generate password")
	}
	db.Create(&models.User{Username: input.Username, Password: string(hashPassword), FullName: input.Fullname})
	return nil
}
