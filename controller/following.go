package controller

import (
	"net/http"
	"regexp"
	"strconv"
	requestinput "tulisaja/request-input/following"
	service "tulisaja/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// Following godoc
//
//	@Summary	Get following user.
//	@Tags		Following
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//
//	@Produce	json
//	@Router		/following [get]
func GetFollowingUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	userMap, _ := c.Get("user")
	userId := int(userMap.(jwt.MapClaims)["id"].(float64))

	listFollowingUser, err := service.GetFollowingUser(db, userId)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Username not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success follow user", "data": listFollowingUser})
	return
}

// Following godoc
//
//	@Summary	Follow a user.
//	@Tags		Following
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		Body	body	requestinput.FollowUserInput	true	"the body to follow a user"
//
//	@Produce	json
//	@Router		/following [post]
func FollowUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input requestinput.FollowUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Please check again your input", "errors": err.Error()})
		return
	}

	userMap, _ := c.Get("user")
	userId := int(userMap.(jwt.MapClaims)["id"].(float64))

	if err := service.FollowUser(db, userId, input.Username); err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Username not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success follow user"})
	return
}

// Following godoc
//
//	@Summary	Delete a following user.
//	@Tags		Following
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//	@Param		id			path	int		true	"Following id"
//
//
//	@Produce	json
//	@Router		/following/{id} [delete]
func DeleteFollowingUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userMap, _ := c.Get("user")
	tmpFollowingId := c.Param("id")
	userId := int(userMap.(jwt.MapClaims)["id"].(float64))

	if match, _ := regexp.MatchString("[0-9]", tmpFollowingId); !match {
		c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "Following id not found"})
		return
	}

	v, _ := strconv.Atoi(tmpFollowingId)
	if err := service.DeleteFollowingUser(db, userId, v); err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success unfollow user"})
	return
}
