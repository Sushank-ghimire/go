package config

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func ConnectDatabase() {
	// Load environment variables
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Get database credentials from environment variables
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open connection to the database
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	fmt.Println("Database connected successfully!")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
