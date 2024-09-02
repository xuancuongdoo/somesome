package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // postgres driver
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByIDs(int) (*Account, error)
	GetAllAccounts() ([]Account, error)
}

type DBConfig struct {
	User string
	Password string
	DBName string
	SSLMode string
}

type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore creates a new PostgrseStore given the DBConfig.
//
// The connection string is built using the config fields.
// The connection string is logged when the function is called.
// The function returns a new PostgrseStore and an error.
// The error is returned if there is an issue opening the db with the
// connection string, or if the db cannot be pinged.
func NewPostgresStore(config DBConfig) (*PostgresStore, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.User, config.Password, config.DBName, config.SSLMode)
	log.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}


func (postgresStore *PostgresStore) Init() error {
	return postgresStore.CreateAccountTable()
}
func (postgresStore *PostgresStore)CreateAccountTable() error{
	creationQuery := `
	create table account if not exists (
	id serial primary key,
	first_name varchar(50),	
	last_name varchar(50),
	number serial,
	balance serial,
	created_at timestamp
	)`
	_, err := postgresStore.db.Exec(creationQuery)
	return err
}


func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName: lastName,
	}
}
func CreateAccount(*Account) error{
	return nil
}

func DeleteAccount(int) error {
	return nil
}
func UpdateAccount(*Account) error {
	return nil
}
func GetAccountByIDs(int) (*Account, error) {
	return nil, nil
}
func GetAllAccounts() ([]Account, error) {
	return nil, nil
}


