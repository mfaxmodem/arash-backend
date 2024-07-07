package models

import (
<<<<<<< HEAD
	"gorm.io/gorm"
=======
	"github.com/jinzhu/gorm"
>>>>>>> 8f20cb1 (edareh)
)

type Role struct {
	gorm.Model
	Name string `gorm:"unique_index;not null"`
}
