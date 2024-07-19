package services

import (
	"arash-website/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"regexp"
)

type UserService struct {
	Db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{Db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Check email structure
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(user.Email) {
		return errors.New("invalid email address")
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 12)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)

	// Validate username and email uniqueness
	if err := s.Db.Where("username = ?", user.Username).First(&models.User{}).Error; err == nil {
		return errors.New("username already exists")
	}
	if err := s.Db.Where("email = ?", user.Email).First(&models.User{}).Error; err == nil {
		return errors.New("email already exists")
	}

	// Create the user
	return s.Db.Create(user).Error
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := s.Db.Find(&users).Error
	return users, err
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := s.Db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 12)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)

	// Validate username and email uniqueness
	if err := s.Db.Where("username = ?", user.Username).First(&models.User{}).Error; err == nil {
		return errors.New("username already exists")
	}
	if err := s.Db.Where("email = ?", user.Email).First(&models.User{}).Error; err == nil {
		return errors.New("email already exists")
	}

	// Update the user
	return s.Db.Save(user).Error
}

func (s *UserService) DeleteUser(id string) error {
	return s.Db.Delete(&models.User{}, "id = ?", id).Error
}

func (s *UserService) VerifyPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}
