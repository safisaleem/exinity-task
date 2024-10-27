package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDatabase() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		return nil, fmt.Errorf("could not load .env file")
	}

	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		return nil, fmt.Errorf("DB_CONNECTION_STRING not set in environment")
	}

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		if err == nil {
			log.Println("Connected to database")
			return db, nil
		}

		log.Println("Failed to connect to database. Retrying...")
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after %d retries", maxRetries)
}
