package services

import (
	"arash-website/models"
	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	if err := s.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := s.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	if err := s.DB.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	if err := s.DB.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := s.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
