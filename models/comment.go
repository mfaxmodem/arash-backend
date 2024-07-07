package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null"`
	ProductID uint
	UserID    uint
}
