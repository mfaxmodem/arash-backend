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
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Product{},
		&models.Category{},
		&models.ProductImage{},
		&models.ProductTag{},
		&models.Tag{},
		&models.Comment{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Apply CSRF middleware
	r.Use(middlewares.AuthMiddleware())

	// Create an API group
	api := r.Group("/api")

	// Setup services
	productService := services.NewProductService(config.DB)
	categoryService := services.NewCategoryService(config.DB)
	commentService := services.NewCommentService(config.DB)
	userService := services.NewUserService(config.DB)
	roleService := services.NewRoleService(config.DB)

	// Setup controllers
	commentController := controllers.NewCommentController(commentService)
	roleController := controllers.NewRoleController(roleService)
	userController := controllers.NewUserController(*userService)

	// Setup routes
	routes.SetupProductRoutes(api, productService)
	routes.SetupCategoryRoutes(api, categoryService)
	routes.SetupCommentRoutes(api, commentController)
	routes.SetupUserRoutes(api, userController)
	routes.SetupRoleRoutes(api, roleController)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
