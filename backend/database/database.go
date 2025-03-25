// sqlite connection string
package database

import (
	"github.com/patrickhcnguyen/PassManager/backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	DB = db

	err = db.AutoMigrate(&models.User{}, &models.PasswordEntry{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
