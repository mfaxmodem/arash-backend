package controllers

import (
	"arash-website/models"
	"arash-website/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContactController struct {
	Service *services.ContactService
}

func NewContactController(service *services.ContactService) *ContactController {
	return &ContactController{Service: service}
}

func (cc *ContactController) SubmitContactForm(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateContact(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.Service.SubmitContactForm(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact form submitted successfully"})
}
