package main

import (
	"arash-website/config"
	"arash-website/models"
	"arash-website/routes"
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
	if err := config.DB.AutoMigrate(&models.User{}, &models.Role{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupUserRoutes(r)
	routes.SetupRoleRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
