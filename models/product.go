package models

<<<<<<< HEAD
import "gorm.io/gorm"
=======
import "github.com/jinzhu/gorm"
>>>>>>> 8f20cb1 (edareh)

type Product struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	CategoryID  uint
}
