package service

import (
	"errors"
	"tulisaja/models"
	requestinput "tulisaja/request-input/article"

	"gorm.io/gorm"
)

type Filters struct {
	Limit int
	Page  int
}

func GetRandomArticles(db *gorm.DB, filters Filters) ([]models.Article, error) {
	var randomArticles []models.Article
	err := db.Select("id", "username").Order("RAND()").Limit(filters.Limit).Offset((filters.Page-1)*filters.Limit).Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "username")
	}).Find(&randomArticles).Error

	if err != nil {
		return nil, errors.New("Internal server error")
	}

	return randomArticles, nil
}

func GetArticles(db *gorm.DB, filters Filters, username string) ([]models.Article, error) {
	var articles []models.Article
	var user models.User

	if err := db.Select("id").Where("username = ?", username).Find(&user).Error; err != nil {
		return nil, errors.New("username not found")
	}

	err := db.Select("title", "content").Where("user_id = ?", user.ID).Limit(filters.Limit).Offset((filters.Page - 1) * filters.Limit).Find(&articles).Error

	if err != nil {
		return nil, errors.New("Internal server error")
	}

	return articles, nil
}

func GetArticle(db *gorm.DB) {

}

func CreateArticle(db *gorm.DB, input requestinput.SaveArticleInput, userId interface{}) error {
	id := uint(userId.(float64))
	return db.Create(&models.Article{Content: input.Content, Title: input.Title, UserId: id}).Error
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

	return db.Model(&article).Where("id = ?", articleId).Update("title", input.Title).Update("content", input.Content).Error
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
