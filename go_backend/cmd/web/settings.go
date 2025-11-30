package main

import (
	"chess_log/go_backend/internal/db"
	"chess_log/go_backend/internal/logger"
	"chess_log/go_backend/internal/models"

	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// loadEnv loads environment variables from a .env file if not running in Docker.
var envPath = "./"

func loadEnv() error {

	os.Setenv("TZ", "UTC")

	// Check for DOCKERIZED (set as env in Docker Compose)
	if os.Getenv("DOCKERIZED") == "true" {
		// In Docker, assume envs are passed by Docker and skip loading .env files.
		return nil
	}

	var filename string
	filename = ".env"

	fullPath := fmt.Sprintf("%s%s", envPath, filename)
	err := godotenv.Load(fullPath)
	if err != nil {
		fmt.Printf("[envConfig] WARNING: Env file does not exist at: %s\n", fullPath)
	}
	return err
}

// ---------------------------------------------------------------------

func initializeDatabase() {
	// Initialize database connection
	dbPool, err := db.OpenDB()
	if err != nil {
		logger.Error("Error connecting to the database", zap.Error(err))
		os.Exit(1)
	} else {
		logger.Info("Database connection established successfully")
	}

	// Assign the database connection to the app configuration
	app.DB = dbPool

	app.Models, err = models.New(dbPool)
	if err != nil {
		logger.Error("Error initializing models", zap.Error(err))
	}
}
