package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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


func getID(request *http.Request) (int, error) {
	idStr := mux.Vars(request)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid account id: %s", idStr)
	}
	return id, nil
}

// WriteJSON writes the given data to the http.ResponseWriter as a JSON
// response with the given status code.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// makeHTTPHandler wraps an apiFunc with error handling. If the apiFunc
// returns an error, makeHTTPHandler will write a JSON response with a
// 400 status code and an ApiError object with the error message.
func makeHTTPHandler(_func apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := _func(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
