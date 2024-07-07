package services

import (
	"arash-website/models"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	if err := s.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}
