package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	CategoryID  uint
}
