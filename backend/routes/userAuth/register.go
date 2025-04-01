// register username, email, and password
// bcrypt for hashing master password

package userAuth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/patrickhcnguyen/PassManager/backend/database"
	"github.com/patrickhcnguyen/PassManager/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// register master user
	var input struct {
		Email          string `json:"email" binding:"required,email"`
		Username       string `json:"username" binding:"required,min=3,max=32"`
		MasterPassword string `json:"master_password" binding:"required,min=8"`
	}

	// try reading user input to JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:          input.Email,
		Username:       input.Username,
		MasterPassword: input.MasterPassword,
	}
	// hash master password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.MasterPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash master password"})
		return
	}
	user.MasterPassword = string(hashedPassword)

	// save user to database
	result := database.DB.Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			c.JSON(http.StatusConflict, gin.H{"error": "A user with this email or username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
