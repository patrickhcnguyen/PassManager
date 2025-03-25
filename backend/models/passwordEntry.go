package models

import "gorm.io/gorm"

type PasswordEntry struct {
	gorm.Model
	UserID            uint   `gorm:"not null"`
	SiteName          string `gorm:"not null"`
	SiteUsername      string `gorm:"not null"`
	EncryptedPassword string `gorm:"not null"`
}
