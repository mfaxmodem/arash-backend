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

func (s *CommentService) GetCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	if err := s.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (s *CommentService) GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	if err := s.DB.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) UpdateComment(comment *models.Comment) error {
	if err := s.DB.Save(comment).Error; err != nil {
		return err
	}
	return nil
}

func (s *CommentService) DeleteComment(id uint) error {
	if err := s.DB.Delete(&models.Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}
