package routes

import (
	"arash-website/controllers"
	"arash-website/services"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.RouterGroup, productService *services.ProductService) {
	productController := controllers.NewProductController(productService)
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/", productController.GetAllProducts)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.PUT("/:id", productController.UpdateProduct)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}
}
