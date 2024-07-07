package models

<<<<<<< HEAD
import "gorm.io/gorm"
=======
import "github.com/jinzhu/gorm"
>>>>>>> 8f20cb1 (edareh)

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null"`
	ProductID uint
	UserID    uint
}
