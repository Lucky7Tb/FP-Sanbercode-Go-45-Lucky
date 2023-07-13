package controller

import (
	"net/http"
	service "tulisaja/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Followers godoc
//
//	@Summary	Get follower user.
//	@Tags		Followers
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//
//	@Produce	json
//	@Router		/followers [get]
func GetFollowers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	userMap, _ := c.Get("user")
	userId := int(userMap.(jwt.MapClaims)["id"].(float64))

	listFollowers, err := service.GetFollowers(db, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success get followers user", "data": listFollowers})
	return
}
