package main

import (
	"arash-website/config"
	"arash-website/controllers"
	"arash-website/middlewares"
	"arash-website/models"
	"arash-website/routes"
	"arash-website/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func main() {

	// Initialize the database connection
	config.InitDB()
	defer func(DB *gorm.DB) {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Could not get database: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Could not close database: %v", err)
		}
	}(config.DB)

	// Migrate the schema
	if err := config.DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Product{}, &models.Category{}, &models.ProductImage{},
		&models.ProductTag{}, &models.Tag{}, &models.Comment{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Apply CSRF middleware
	r.Use(middlewares.AuthMiddleware())

	// Setup routes
	publicRoutes(r)
	protectedRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

func publicRoutes(r *gin.Engine) {
	routes.SetupUserRoutes(r)
	routes.SetupRoleRoutes(r)
	routes.SetupProductRoutes(r, services.NewProductService(config.DB))
	routes.SetupCategoryRoutes(r, services.NewCategoryService(config.DB))
	routes.SetupCommentRoutes(r, controllers.NewCommentController(services.NewCommentService(config.DB)))
}

func protectedRoutes(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		commentController := controllers.NewCommentController(services.NewCommentService(config.DB))
		protected.POST("/comments/approve/:id", commentController.ApproveComment)
		// Define more protected routes here
	}
}
