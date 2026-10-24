package models

import (
	"time"
)

// Repository interface for career operations
type CareerRepository interface {
	FindAll() ([]Career, error)
	FindByID(id string) (*Career, error)
	FindBySlug(slug string) (*Career, error)
	FindByCategory(category string) ([]Career, error)
	FindByTags(tags []string) ([]Career, error)
	Create(career *Career) error
	Update(career *Career) error
	Delete(id string) error
	Count() (int64, error)
}

// DatabaseConfig holds MongoDB connection configuration
type DatabaseConfig struct {
	Host         string `mapstructure:"MONGODB_HOST"`
	Port         string `mapstructure:"MONGODB_PORT"`
	Name         string `mapstructure:"MONGODB_NAME"`
	Username     string `mapstructure:"MONGODB_USERNAME"`
	Password     string `mapstructure:"MONGODB_PASSWORD"`
	ConnectionStr string `mapstructure:"MONGODB_URI"`
}

// Config holds application configuration
type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string `mapstructure:"PORT"`
}

// Response represents API response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// PaginatedResponse represents paginated API response
type PaginatedResponse struct {
	Items      []Career `json:"items"`
	Total      int64    `json:"total"`
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	TotalPages int      `json:"total_pages"`
}

// CareerStore provides career storage interface
type CareerStore interface {
	FindAll() ([]Career, error)
	FindByID(id string) (*Career, error)
	FindBySlug(slug string) (*Career, error)
	FindByCategory(category string) ([]Career, error)
	FindByTags(tags []string) ([]Career, error)
	Create(career *Career) error
	Update(career *Career) error
	Delete(id string) error
	Count() (int64, error)
}

// SeedData represents pre-seed data for careers
type SeedData struct {
	Careers  []Career   `json:"careers"`
	SeededAt time.Time  `json:"seeded_at"`
}
