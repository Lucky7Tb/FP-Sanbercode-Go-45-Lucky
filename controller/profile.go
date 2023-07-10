package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	requestinput "tulisaja/request-input/profile"
	service "tulisaja/service"
)

// Profile godoc
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Tags		Profile
//	@Produce	json
//	@Router		/profile [GET]
func GetProfile(c *gin.Context) {
	userMap, _ := c.Get("user")

	user := struct {
		FullName string `json:"full_name"`
		Username string `json:"user_name"`
	}{
		FullName: userMap.(jwt.MapClaims)["full_name"].(string),
		Username: userMap.(jwt.MapClaims)["username"].(string),
	}
	fmt.Println(userMap.(jwt.MapClaims)["full_name"].(string))
	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "message": "success get profile", "data": user})
}

// Change Password godoc
//
//	@Summary	Change password user.
//	@Tags		Profile
//
//	@Param		Authorization	header	string	true	"Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
//	@Security	BearerToken
//
//	@Param		Body	body	requestinput.ChangePasswordInput	true	"the body to change password"
//
//	@Produce	json
//	@Router		/profile/change-password [post]
func ChangePassword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input requestinput.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Please check again your input", "errors": err.Error()})
		return
	}

	userMap, _ := c.Get("user")

	if err := service.ChangePassword(db, input, userMap.(jwt.MapClaims)["id"]); err != nil {
		switch err.Error() {
		case "record not found":
			c.JSON(http.StatusNotFound, gin.H{"status_code": http.StatusNotFound, "messsage": "User not found"})
			return
		case "old password error":
			c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Old password is wrong"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Failed to change password"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success change password"})
	return
}
