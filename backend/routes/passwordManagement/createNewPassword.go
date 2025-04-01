package passwordManagement

import (
	"net/http"
	"os"
	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/patrickhcnguyen/PassManager/backend/database"
	"github.com/patrickhcnguyen/PassManager/backend/models"
	"crypto/aes"
	"encoding/base64"
	"bytes"
	"crypto/cipher"
	"log"
	"fmt"
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
	// encrypt password
	encryptedPassword, err := EncryptPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to encrypt password"})
		return
	}
	password.EncryptedPassword = encryptedPassword

	database.DB.Create(&password)

	c.JSON(http.StatusOK, gin.H{"message": "password created successfully"})
}

func EncryptPassword(password string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := []byte(os.Getenv("key"))
	iv := []byte(os.Getenv("iv"))

	// Ensure IV is exactly 16 bytes
	if len(iv) != aes.BlockSize {
		return "", fmt.Errorf("IV must be exactly %d bytes", aes.BlockSize)
	}

	// Ensure key is 16, 24, or 32 bytes (for AES-128, AES-192, or AES-256)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("key must be 16, 24, or 32 bytes")
	}

	// init cipher using key
	var plainTextBlock []byte
	length := len(password)

	if length % 16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	
	// Create ciphertext buffer with same length as plaintext
    ciphertext := make([]byte, len(plainTextBlock))

	// create a byte array of the plain text's length and start encrypting
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainTextBlock)

	// convert to base64
	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil	
}
