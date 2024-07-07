package models

<<<<<<< HEAD
import "gorm.io/gorm"
=======
import "github.com/jinzhu/gorm"
>>>>>>> 8f20cb1 (edareh)

type Contact struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"not null"`
	Message string `gorm:"not null"`
}
