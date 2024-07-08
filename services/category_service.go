package services

import (
	"arash-website/models"
	"fmt"
	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	var existingCategory models.Category
	if err := s.DB.Where("name = ?", category.Name).First(&existingCategory).Error; err == nil {
		return fmt.Errorf("category with name '%s' already exists", category.Name)
	}

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

func (s *CategoryService) DeleteCategory(category *models.Category) error {
	if err := s.DB.Delete(category).Error; err != nil {
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
