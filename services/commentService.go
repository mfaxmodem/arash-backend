package services

import (
	"arash-website/models"
	"gorm.io/gorm"
)

type CommentService struct {
	DB *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{DB: db}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	if err := s.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (s *CommentService) ApproveComment(id uint) error {
	if err := s.DB.Model(&models.Comment{}).Where("id = ?", id).Update("approved", true).Error; err != nil {
		return err
	}
	return nil
}

func (s *CommentService) GetCommentsByProductID(productID uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := s.DB.Where("product_id = ? AND approved = ?", productID, true).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	if err := s.DB.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
