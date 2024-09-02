package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() { 
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := DBConfig{
        User:     "postgres",
        Password: "test",
        DBName:   "go-postgres",
        SSLMode:  "disable",
    }

	
	store, err := NewPostgresStore(config)
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store)
	server.Run()
}