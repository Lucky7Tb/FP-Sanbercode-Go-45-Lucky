package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Profile godoc
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
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

func ChangePassword(c *gin.Context) {

}
