package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/hari7k-creator/my-career-api/internal/config"
	carrier "github.com/hari7k-creator/my-career-api/internal/handlers"
	"github.com/hari7k-creator/my-career-api/internal/middleware"
	"github.com/hari7k-creator/my-career-api/internal/models"
	"github.com/hari7k-creator/my-career-api/internal/repository"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	// Validator initialization handled in models package
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize MongoDB client
	client, err := connectDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test connection
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to MongoDB successfully!")

	// Initialize repository
	repo := repository.NewCareerRepository(client)

	// Initialize database (create indexes, etc.)
	if err := repository.InitializeDatabase(client); err != nil {
		log.Printf("Warning: Failed to initialize database: %v", err)
	}

	// Initialize router
	r := mux.NewRouter()

	// Add CORS middleware
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	r.Use(cors)
	r.Use(middleware.LoggerMiddleware)

	// Initialize handlers
	careerHandler := carrier.NewCareerHandler(repo)

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Career routes
	api.HandleFunc("/careers", careerHandler.GetCareers).Methods("GET")
	api.HandleFunc("/careers", careerHandler.CreateCareer).Methods("POST")
	api.HandleFunc("/careers/slug/{slug}", careerHandler.GetCareerBySlug).Methods("GET")
	api.HandleFunc("/careers/{id}", careerHandler.GetCareerByID).Methods("GET")
	api.HandleFunc("/careers/{id}", careerHandler.UpdateCareer).Methods("PUT")
	api.HandleFunc("/careers/{id}", careerHandler.DeleteCareer).Methods("DELETE")
	api.HandleFunc("/categories", careerHandler.GetCategories).Methods("GET")
	api.HandleFunc("/difficulties", careerHandler.GetDifficulties).Methods("GET")

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "healthy",
			"message": "Career Architecture API is running",
		})
	}).Methods("GET")

	// Start server
	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Starting server on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func connectDatabase(dbConfig models.DatabaseConfig) (*mongo.Client, error) {
	var clientOptions *options.ClientOptions

	if dbConfig.ConnectionStr != "" {
		clientOptions = options.Client().ApplyURI(dbConfig.ConnectionStr)
	} else {
		clientOptions = options.Client().ApplyURI(
			formatMongoURI(dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name),
		)
	}

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func formatMongoURI(host, port, username, password, dbName string) string {
	if username != "" && password != "" {
		return "mongodb://" + username + ":" + password + "@" + host + ":" + port + "/" + dbName
	}
	return "mongodb://" + host + ":" + port + "/" + dbName
}
