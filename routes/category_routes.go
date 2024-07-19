package routes

import (
	"arash-website/controllers"
	"arash-website/services"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.RouterGroup, categoryService *services.CategoryService) {
	categoryController := controllers.NewCategoryController(categoryService)
	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.GET("/", categoryController.GetAllCategories)
		categoryRoutes.GET("/:id", categoryController.GetCategoryByID)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}
}
