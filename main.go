package main

import (
	"log"
)

func main() {
	// Load environment variables
	if err := loadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database configuration
	config := getDBConfig()

	// Create and initialize Postgres store
	store, err := NewPostgresStore(config)
	if err != nil {
		log.Fatalf("Failed to create Postgres store: %v", err)
	}

	if err := store.Init(); err != nil {
		log.Fatalf("Failed to initialize store: %v", err)
	}

	// Start API server
	server := NewAPIServer(":3000", store)
	server.Run()
}
