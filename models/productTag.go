package models

import "gorm.io/gorm"

type ProductTag struct {
	gorm.Model
	ProductID uint
	TagID     uint
}
