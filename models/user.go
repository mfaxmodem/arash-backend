package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	//Password     string `json:"-"`
	PasswordHash string `json:"-"`
}

// SetPassword متدی است که پسورد را هش می‌کند و در فیلد PasswordHash ذخیره می‌کند.
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

// CheckPassword متدی است که پسورد ورودی را با هش ذخیره شده در دیتابیس مقایسه می‌کند.
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}
