package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NewAccount creates a new Account given the first name and last name.
// The Account's ID is a random number between 0 and 99,999.
// The Account's Number is a random number between 0 and 99,999,999.
// NewAPIServer creates a new APIServer given the listen address and the store.
// The APIServer is the JSON API server that listens on the given address.
// The store is the underlying storage for the API server.
func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

// Run starts the API server listening on s.listenAddr.
// It registers a single route: GET /account/{account_id}.
func (apiServer *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/accounts", makeHTTPHandler(apiServer.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandler(apiServer.handleGetAccountByID))
	router.HandleFunc("/account", makeHTTPHandler(apiServer.handleCreateAccount))
	router.HandleFunc("/transfer", makeHTTPHandler(apiServer.handleTransferAccount))
	log.Println("JSON API server on port : ", apiServer.listenAddr)
	http.ListenAndServe(apiServer.listenAddr, router)
}

func (apiServer *APIServer) handleAccount(response http.ResponseWriter, request *http.Request) error {
	if request.Method == "GET" {
		return apiServer.handleGetAccount(response, request)
	}
	if request.Method == "DELETE" {
		return apiServer.handleDeleteAccount(response, request)
	}
	if request.Method == "POST" {
		return apiServer.handleCreateAccount(response, request)
	}

	return nil
}

func (apiServer *APIServer) handleGetAccount(response http.ResponseWriter, request *http.Request) error {
	accounts, err := apiServer.store.GetAllAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(response, http.StatusOK, accounts)
}

func (apiServer *APIServer) handleGetAccountByID(response http.ResponseWriter, request *http.Request) error {
	fmt.Println(request.Method)
	if request.Method == "GET" {
		idStr := mux.Vars(request)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return fmt.Errorf("invalid account id: %s", idStr)
		}
		account, err := apiServer.store.GetAccountByID(id)
		if err != nil {
			return err
		}
		return WriteJSON(response, http.StatusOK, account)
	}
	if request.Method == "DELETE" {
		return apiServer.handleDeleteAccount(response, request)
	}
	return fmt.Errorf("invalid method: %s", request.Method)
}
func (apiServer *APIServer) handleCreateAccount(response http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(CreateAccountRequest)

	if errors := json.NewDecoder(r.Body).Decode(createAccountRequest); errors != nil {
		return errors
	}
	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName)
	if errors := apiServer.store.CreateAccount(account); errors != nil {
		return errors
	}
	return WriteJSON(response, http.StatusOK, createAccountRequest)
}

func (apiServer *APIServer) handleDeleteAccount(response http.ResponseWriter, request *http.Request) error {
	id, err := getID(request)
	log.Println("Deleting account with ID:", id)
	if err != nil {
		return err
	}
	if err := apiServer.store.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJSON(response, http.StatusOK, map[string]int{"deleted": id})
}
func (apiServer *APIServer) handleTransferAccount(response http.ResponseWriter, request *http.Request) error {
	transferRequest := new(TransferRequest)

	if err := json.NewDecoder(request.Body).Decode(transferRequest); err != nil {
		return nil
	}

	defer request.Body.Close()
	return WriteJSON(response, http.StatusOK, transferRequest)
}
