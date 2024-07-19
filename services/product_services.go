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

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := s.DB.Preload("Images").Preload("Tags").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductService) UpdateProduct(id uint, updatedProduct *models.Product) (*models.Product, error) {
	var product models.Product
	if err := s.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	if err := s.DB.Model(&product).Updates(updatedProduct).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	if err := s.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := s.DB.Preload("Images").Preload("Tags").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
