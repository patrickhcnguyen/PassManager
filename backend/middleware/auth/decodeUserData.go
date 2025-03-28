package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// extracts user data from jwt token from claims and returns it as a json

func UserHandler(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if claims, ok := claims.(jwt.MapClaims); ok {
		userData := gin.H{
			"username": claims["username"],
			"email":    claims["email"],
		}
		c.JSON(200, userData)
	} else {
		c.JSON(401, gin.H{"error": "Unauthorized"})
	}
}
