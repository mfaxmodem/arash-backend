package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	CategoryID  uint
	Images      []ProductImage `gorm:"foreignKey:ProductID"`
	Tags        []Tag          `gorm:"many2many:product_tags;"`
	Slug        string         `gorm:"uniqueIndex"`
}
