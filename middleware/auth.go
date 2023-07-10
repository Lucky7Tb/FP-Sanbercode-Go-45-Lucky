package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func VerifyJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(strings.Split(authHeader, " ")) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"status_code": http.StatusUnauthorized, "message": "Jwt key is not valid"})
			c.Abort()
			return
		}

		token := strings.Split(authHeader, " ")[1]

		secretKey, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status_code": http.StatusUnauthorized, "message": "Jwt key is not valid"})
			c.Abort()
			return
		}

		userData, ok := secretKey.Claims.(jwt.MapClaims)
		if !ok && !secretKey.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"status_code": http.StatusUnauthorized, "message": "Jwt key is not valid"})
			c.Abort()
			return
		}

		c.Set("user", userData)
		c.Next()
	}
}
