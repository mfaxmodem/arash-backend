package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Slug string `gorm:"uniqueIndex"`
}
