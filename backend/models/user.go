package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID       string    `gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `gorm:"unique"`
	Picture      string    `json:"picture"`
	Provider     string    `json:"provider" gorm:"default:'google'"`
	LastLogin    time.Time `json:"last_login"`
	RefreshToken string    `json:"-"`
}

type TempUser struct {
	gorm.Model
	TempID        string `gorm:"primaryKey"`
	BrowserID     string `gorm:"index"`
	IPAddress     string
	LastActive    time.Time
	ConvertedToID *string `gorm:"index"` // Points to User.UserID after conversion
}
