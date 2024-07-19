package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	AuthorName string `gorm:"not null"`
	Email      string
	Content    string `gorm:"not null"`
	Approved   bool   `gorm:"default:false"`
	ProductID  uint
}
