package controller

import (
	"net/http"
	"regexp"
	"strconv"
	requestinput "tulisaja/request-input/article"
	service "tulisaja/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Get Article godoc
//
//	@Summary	Get random article.
//	@Tags		Articles
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Produce	json
//	@Router		/articles [get]
func GetRandomArticles(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	articles, err := service.GetRandomArticles(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Failed get articles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success to create article", "data": articles})
	return
}

// Get user article godoc
//
//	@Summary	Get user article.
//	@Tags		Articles
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		limit		query	int		false	"Limit returning value"
//	@Param		page		query	int		false	"Paging"
//	@Param		username	path	string	true	"user username"
//	@Produce	json
//	@Router		/articles/{username} [get]
func GetArticles(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var limit int = 10
	var page int = 1
	var username = c.Param("username")

	if match, _ := regexp.MatchString("[0-9]", strconv.Itoa(limit)); c.Query("limit") != "" && match {
		v, _ := strconv.Atoi(c.Query("limit"))
		limit = v
	}

	if match, _ := regexp.MatchString("[0-9]", strconv.Itoa(page)); c.Query("page") != "" && match {
		v, _ := strconv.Atoi(c.Query("page"))
		page = v
	}

	filters := service.Filters{
		Limit: limit,
		Page:  page,
	}

	articles, err := service.GetArticles(db, filters, username)
	if err != nil {
		switch err.Error() {
		case "username not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Username not found"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Internal server error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success to get article", "data": articles})
	return
}

// Get detail article godoc
//
//	@Summary	Get detail article.
//	@Tags		Articles
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		username	path	string	true	"user username"
//	@Param		id			path	int		true	"user username"
//	@Produce	json
//	@Router		/articles/{username}/{id} [get]
func GetArticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var username = c.Param("username")
	var articleId = c.Param("id")

	if match, _ := regexp.MatchString("[0-9]", articleId); !match {
		c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article not found"})
		return
	}

	v, err := strconv.Atoi(articleId)
	article, err := service.GetArticle(db, username, v)
	if err != nil {
		switch err.Error() {
		case "not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Username not found"})
			return
		case "record not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article not found"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Internal server error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success to get article", "data": article})
	return
}

// Create Article godoc
//
//	@Summary	Create article.
//	@Tags		Articles
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		Body	body	requestinput.SaveArticleInput	true	"the body to create article"
//
//	@Produce	json
//	@Router		/articles [post]
func CreateArticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input requestinput.SaveArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Please check again your input", "errors": err.Error()})
		return
	}

	userMap, _ := c.Get("user")

	if err := service.CreateArticle(db, input, userMap.(jwt.MapClaims)["id"]); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Failed to create article", "errors": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status_code": http.StatusCreated, "messsage": "Success to create article"})
	return
}

// Update Article godoc
//
//	@Summary	Update article.
//	@Tags		Articles
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		Body	body	requestinput.SaveArticleInput	true	"the body to update article"
//	@Param		id		path	int								true	"Article id"
//
//	@Produce	json
//	@Router		/articles/{id} [put]
func UpdateArticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input requestinput.SaveArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Please check again your input", "errors": err.Error()})
		return
	}

	userMap, _ := c.Get("user")
	articleId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article id not found", "errors": err.Error()})
		return
	}

	if err := service.UpdateArticle(db, input, articleId, userMap.(jwt.MapClaims)["id"]); err != nil {
		switch err.Error() {
		case "record not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article not found"})
			return
		case "forbidden update":
			c.JSON(http.StatusForbidden, gin.H{"status_code": http.StatusForbidden, "messsage": "You don't have permission to update this article"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Failed to update article"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success to update article"})
	return
}

// Delete Article godoc
//
//	@Summary	Delete article.
//	@Tags		Articles
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		id	path	int	true	"Article id"
//
//	@Produce	json
//	@Router		/articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userMap, _ := c.Get("user")
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Article id not found", "errors": err.Error()})
		return
	}

	if err := service.DeleteArticle(db, articleId, userMap.(jwt.MapClaims)["id"]); err != nil {
		switch err.Error() {
		case "record not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article not found"})
			return
		case "forbidden delete":
			c.JSON(http.StatusForbidden, gin.H{"status_code": http.StatusForbidden, "messsage": "You don't have permission to delete this article"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Failed to delete article"})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
	return
}

// Comment Article godoc
//
//	@Summary	Comment article.
//	@Tags		Articles
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		username	path	string						true	"user username"
//	@Param		id			path	int							true	"Article id"
//
//	@Param		Body		body	requestinput.CommentInput	true	"the body to comment an article"
//	@Produce	json
//	@Router		/articles/{username}/{id}/comment [post]
func CommentArticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var username = c.Param("username")
	var tmpArticleId = c.Param("id")
	userMap, _ := c.Get("user")

	if match, _ := regexp.MatchString("[0-9]", tmpArticleId); !match {
		c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article not found"})
		return
	}

	var input requestinput.CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Please check again your input", "errors": err.Error()})
		return
	}

	artileId, _ := strconv.Atoi(tmpArticleId)
	input.UserID = int(userMap.(jwt.MapClaims)["id"].(float64))
	if err := service.CommentArticle(db, artileId, username, input); err != nil {
		switch err.Error() {
		case "not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Username not found"})
			return
		case "record not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article not found"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Failed"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success command on article"})
	return
}

// Likes Article godoc
//
//	@Summary	Likes article.
//	@Tags		Articles
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		username	path	string	true	"user username"
//	@Param		id			path	int		true	"Article id"
//
//	@Produce	json
//	@Router		/articles/{username}/{id}/like [post]
func LikeArticle(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var username = c.Param("username")
	var tmpArticleId = c.Param("id")

	if match, _ := regexp.MatchString("[0-9]", tmpArticleId); !match {
		c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article not found"})
		return
	}

	artileId, _ := strconv.Atoi(tmpArticleId)
	if err := service.LikeArticle(db, artileId, username); err != nil {
		switch err.Error() {
		case "not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Username not found"})
			return
		case "record not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Article not found"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Failed"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success like on article"})
	return
}
