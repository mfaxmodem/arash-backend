package routes

import (
	"arash-website/controllers"
	"arash-website/services"
	"github.com/gin-gonic/gin"
)

func SetupContactRoutes(r *gin.Engine, contactService *services.ContactService) {
	controller := controllers.NewContactController(contactService)

	r.POST("/contact", controller.SubmitContactForm)
}
