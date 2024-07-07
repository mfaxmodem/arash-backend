package models

<<<<<<< HEAD
import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
=======
import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
>>>>>>> 8f20cb1 (edareh)
}
