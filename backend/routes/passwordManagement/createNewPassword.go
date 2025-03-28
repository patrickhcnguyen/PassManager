package passwordManagement

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/patrickhcnguyen/PassManager/backend/database"
	"github.com/patrickhcnguyen/PassManager/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewPassword(c *gin.Context) {
	var input struct {
		WebsiteName string `json:"website_name" binding:"required"`
		Username    string `json:"username" binding:"required"` // site username
		Password    string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims := c.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	user := models.User{}
	database.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	password := models.PasswordEntry{
		UserID:            user.ID,
		SiteUsername:      input.Username,
		SiteName:          input.WebsiteName,
		EncryptedPassword: input.Password,
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
		return
	}
	password.EncryptedPassword = string(hashedPassword)

	database.DB.Create(&password)

	c.JSON(http.StatusOK, gin.H{"message": "password created successfully"})

}
