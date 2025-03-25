package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string          `gorm:"unique;not null"`
	Username        string          `gorm:"unique;not null"`
	MasterPassword  string          `gorm:"not null"`
	PasswordEntries []PasswordEntry `gorm:"foreignKey:UserID"`
}
