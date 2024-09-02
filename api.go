package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// NewAccount creates a new Account given the first name and last name.
//
// The Account's ID is a random number between 0 and 99,999.
// The Account's Number is a random number between 0 and 99,999,999.

// func NewAccount(firstName string, lastName string) *Account {
// 	return &Account{
// 		ID: rand.Intn(100000),
// 		FirstName: firstName,
// 		LastName: lastName,
// 		Number: int64(rand.Intn(100000000)),
// 	}
// }

// NewAPIServer creates a new APIServer given the listen address and the store.
// The APIServer is the JSON API server that listens on the given address.
// The store is the underlying storage for the API server.
func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr : listenAddr,
		store : store,
	}
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
func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {	
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// Run starts the API server listening on s.listenAddr.
// It registers a single route: GET /account/{account_id}.
func (apiServer *APIServer) Run() {
	router := mux.NewRouter()
    router.HandleFunc("/account", makeHTTPHandler(apiServer.handleAccount))
	log.Println("JSON API server on port : ", apiServer.listenAddr)
	router.HandleFunc("/account/{id}", makeHTTPHandler(apiServer.handleGetAccount))
	http.ListenAndServe(apiServer.listenAddr, router)
}

func (apiServer *APIServer) handleAccount (response http.ResponseWriter, request *http.Request) error {
	if request.Method == "GET" {
		return apiServer.handleGetAccount(response, request)
	}
	if request.Method	== "DELETE" {
		return apiServer.handleDeleteAccount(response, request)
	}
	if request.Method	== "POST" {
		return apiServer.handleCreateAccount(response, request)
	}

	return nil
}

func (apiServer *APIServer) handleGetAccount (response http.ResponseWriter, request *http.Request) error {
	// account := NewAccount("Xuan", "Cuong")
	vars := mux.Vars(request)
	return WriteJSON(response, http.StatusOK, vars)
}
func (apiServer *APIServer) handleCreateAccount (response http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(CreateAccountRequest)
	if errors := json.NewDecoder(r.Body).Decode(createAccountRequest); errors != nil {return errors}
	return WriteJSON(response, http.StatusOK, createAccountRequest)
}

func (apiServer *APIServer) handleDeleteAccount (response http.ResponseWriter, r *http.Request) error {
	return nil
}
func (apiServer *APIServer) handleTransferAccount (response http.ResponseWriter, r *http.Request) error {
	return nil
}
