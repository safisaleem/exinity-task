package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDatabase() (*gorm.DB, error) {
	connectionString := "postgres://postgres:changeme@localhost:5432/exinitytask?sslmode=disable"

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
