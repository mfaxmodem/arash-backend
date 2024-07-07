package routes

import (
	"arash-website/controllers"
	"arash-website/services"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, productService *services.ProductService) {
	productController := controllers.NewProductController(productService)

	router.POST("/products", productController.CreateProduct)
}
