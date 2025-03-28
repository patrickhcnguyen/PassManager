package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
			c.Abort()
			return
		}

		// parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if err != nil {
			c.Abort()
			return
		}

		if !token.Valid {
			c.Abort()
			return
		}

		// set token claims to context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Abort()
			return
		}

		log.Printf("Claims found: %v", claims)
		c.Set("claims", claims)
		c.Next()
	}
}
