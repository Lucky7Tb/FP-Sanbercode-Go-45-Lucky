package controller

import (
	"net/http"
	requestinput "tulisaja/request-input/auth"
	"tulisaja/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type any interface{}

// Login godoc
//
//	@Summary	Login.
//	@Tags		Auth
//
//	@Param		Body	body	requestinput.LoginInput	true	"the body to login"
//
//	@Produce	json
//	@Router		/auth/login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input requestinput.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Please check again your input", "errors": err.Error()})
		return
	}

	token, err := service.Login(db, input)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Username or password not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success login", "data": token})
	return
}

// Register godoc
//
//	@Summary	Register.
//	@Tags		Auth
//
//	@Param		Body	body	requestinput.RegisterInput	true	"the body to register"
//
//	@Produce	json
//	@Router		/auth/register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input requestinput.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": http.StatusBadRequest, "messsage": "Please check again your input", "errors": err.Error()})
		return
	}

	if err := service.Register(db, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status_code": http.StatusInternalServerError, "messsage": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status_code": http.StatusOK, "messsage": "Success register"})
	return
}
