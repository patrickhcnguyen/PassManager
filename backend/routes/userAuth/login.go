// login with email and master password
// remove logs later

package userAuth

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/patrickhcnguyen/PassManager/backend/database"
	"github.com/patrickhcnguyen/PassManager/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) { // checking user credentials and generating jwt token
	var input struct {
		Login          string `json:"login" binding:"required"` // can be email or username
		MasterPassword string `json:"master_password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Login attempt with: %s", input.Login)

	user := models.User{}
	result := database.DB.Where("email = ?", input.Login).First(&user)

	if result.Error != nil {
		log.Printf("Email not found, trying username")
		result = database.DB.Where("username = ?", input.Login).First(&user)
		if result.Error != nil {
			log.Printf("User not found with login: %s", input.Login)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
	}

	log.Printf("Found user: %s", user.Username)

	// compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.MasterPassword), []byte(input.MasterPassword))
	if err != nil {
		log.Printf("Password comparison failed: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // token expiration time
	})

	// sign and get complete encoded token as a string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// set jwt token in cookies
	c.SetCookie("jwt", tokenString, 3600, "/", "localhost", false, true)
	c.SetSameSite(http.SameSiteStrictMode)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

	// send to middleware to check if token is valid
}
