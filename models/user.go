package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique_index;not null"`
	RoleID   uint   // کلید خارجی برای نقش
}
