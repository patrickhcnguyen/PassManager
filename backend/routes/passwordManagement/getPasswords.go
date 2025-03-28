package passwordManagement

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/patrickhcnguyen/PassManager/backend/database"
	"github.com/patrickhcnguyen/PassManager/backend/models"
)

func GetPasswords(c *gin.Context) {
	// get username from claims we handled in middleware
	claims := c.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	// get user's passwords
	var passwords []models.PasswordEntry
	if err := database.DB.
		Joins("JOIN users ON users.id = password_entries.user_id").
		Where("users.username = ?", username).
		Find(&passwords).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch passwords"})
		return
	}

	c.JSON(200, passwords)
}
