package main

import (
	"os"

	"github.com/joho/godotenv"
)

// loadEnv loads environment variables from the .env file
func loadEnv() error {
	return godotenv.Load()
}

// getDBConfig returns the database configuration
func getDBConfig() DBConfig {
	return DBConfig{
		User:     os.Getenv("POSTGRES_USER"),     // Get from environment variable
		Password: os.Getenv("POSTGRES_PASSWORD"), // Get from environment variable
		DBName:   os.Getenv("POSTGRES_NAME"),     // Get from environment variable
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),  // Get from environment variable
	}
}