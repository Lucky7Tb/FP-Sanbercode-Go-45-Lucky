package service

import (
	"errors"
	"tulisaja/models"
	requestinput "tulisaja/request-input/article"

	"gorm.io/gorm"
)

func GetRandomArticles(db *gorm.DB) {

}

func GetArticles(db *gorm.DB) {

}

func GetArticle(db *gorm.DB) {

}

func CreateArticle(db *gorm.DB, input requestinput.SaveArticleInput, userId interface{}) error {
	id := uint(userId.(float64))
	return db.Create(&models.Article{Content: input.Content, UserId: id}).Error
}

func UpdateArticle(db *gorm.DB, input requestinput.SaveArticleInput, articleId int, userId interface{}) error {
	var article models.Article
	result := db.Where("id = ?", articleId).First(&article)

	if result.Error != nil {
		return result.Error
	}

	var id = uint(userId.(float64))
	if article.UserId != id {
		return errors.New("forbidden update")
	}

	return db.Model(&article).Where("id = ?", articleId).Update("content", input.Content).Error
}

func DeleteArticle(db *gorm.DB, articleId int, userId interface{}) error {
	var article models.Article
	result := db.Where("id = ?", articleId).First(&article)

	if result.Error != nil {
		return result.Error
	}

	var id = uint(userId.(float64))
	if article.UserId != id {
		return errors.New("forbidden delete")
	}

	return db.Delete(&article).Error
}

func CommentArticle(db *gorm.DB) {

}

func LikeArticle(db *gorm.DB) {

}
