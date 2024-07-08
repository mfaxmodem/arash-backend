package models

import (
	"github.com/go-playground/validator/v10"
)

type Contact struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Email   string `form:"email" json:"email" binding:"required,email"`
	Phone   string `form:"phone" json:"phone"`
	Message string `form:"message" json:"message" binding:"required"`
}

// ValidateContact validates the Contact struct
func ValidateContact(c *Contact) error {
	validate := validator.New()
	return validate.Struct(c)
}
