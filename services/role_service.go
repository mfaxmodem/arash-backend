package services

import (
	"arash-website/models"
	"gorm.io/gorm"
)

type RoleService struct {
	DB *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{DB: db}
}

func (s *RoleService) CreateRole(role *models.Role) error {
	if err := s.DB.Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	if err := s.DB.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := s.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *RoleService) UpdateRole(role *models.Role) error {
	if err := s.DB.Save(role).Error; err != nil {
		return err
	}
	return nil
}

func (s *RoleService) DeleteRole(id uint) error {
	if err := s.DB.Delete(&models.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}
