package controller

import (
	"net/http"
	"strconv"
	requestinput "tulisaja/request-input/article"
	service "tulisaja/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func GetRandomArticles(c *gin.Context) {

}

func GetArticles(c *gin.Context) {

}

func GetArticle(c *gin.Context) {

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

func CommentArticle(c *gin.Context) {

}

func LikeArticle(c *gin.Context) {

}
