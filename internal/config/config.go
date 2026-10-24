package config

import (
	"fmt"
	"os"

	"career-arch-backend/internal/models"

	"github.com/joho/godotenv"
)

// LoadConfig loads configuration from environment and .env file
func LoadConfig() (*models.Config, error) {
	// Try to load .env file if it exists
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	cfg := &models.Config{
		Server: models.ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
		Database: models.DatabaseConfig{
			Host:         getEnv("MONGODB_HOST", "localhost"),
			Port:         getEnv("MONGODB_PORT", "27017"),
			Name:         getEnv("MONGODB_NAME", "career_arch"),
			Username:     getEnv("MONGODB_USERNAME", ""),
			Password:     getEnv("MONGODB_PASSWORD", ""),
			ConnectionStr: getEnv("MONGODB_URI", ""),
		},
	}

	return cfg, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
