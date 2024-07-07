package models

import (
<<<<<<< HEAD
	"gorm.io/gorm"
=======
	"github.com/jinzhu/gorm"
>>>>>>> 8f20cb1 (edareh)
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique_index;not null"`
	RoleID   uint   // کلید خارجی برای نقش
}
