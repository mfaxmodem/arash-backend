package services

import (
	"arash-website/models"
)

type ContactService struct{}

func NewContactService() *ContactService {
	return &ContactService{}
}

func (cs *ContactService) SubmitContactForm(contact *models.Contact) error {
	// You can implement your logic here, e.g., saving to database, sending emails, etc.
	// For simplicity, we are not implementing database interactions here.
	return nil
}
