package main

import (
	"net/http"
	"time"
)

type APIServer struct {
	listenAddr string
	store Storage
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`	
}

type Account struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Number int64 `json:"number"`
	Balance int64 `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
} 

