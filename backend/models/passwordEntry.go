package models

import "gorm.io/gorm"

type PasswordEntry struct {
	gorm.Model
	User               User   `gorm:"foreignKey:UserID"`
	UserID             uint   `gorm:"not null"`
	SiteName           string `gorm:"not null"`
	SiteUsername       string `gorm:"unique; not null"`
	EncryiptedPassword string `gorm:"unique; not null"`
}
