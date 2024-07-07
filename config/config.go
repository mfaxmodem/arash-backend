package config

import (
<<<<<<< HEAD
	"arash-website/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
=======
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
>>>>>>> 8f20cb1 (edareh)
)

var DB *gorm.DB

func InitDB() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	var err error
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

<<<<<<< HEAD
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", dbHost, dbUser, dbPassword, dbName, dbPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Auto Migrate
	err = DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Product{})
	if err != nil {
		return
	} // Example: Auto migrate all necessary models here
=======
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
>>>>>>> 8f20cb1 (edareh)
	log.Println("Connected to database")
}
