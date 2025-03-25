package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string          `gorm:"unique;not null"`
	MasterPassWord  string          `gorm:"not null"`
	PasswordEntries []PasswordEntry `gorm:"foreignKey:userID"`
}
