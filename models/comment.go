package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null"`
	ProductID uint
	UserID    uint
}
