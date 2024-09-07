package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // postgres driver
	"golang.org/x/exp/rand"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByIDs(int) (*Account, error)
	GetAllAccounts() ([]*Account, error)
}

type DBConfig struct {
	User     string
	Password string
	DBName   string
	SSLMode  string
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

func (postgresStore *PostgresStore) CreateAccountTable() error {
	creationQuery := `
		CREATE TABLE IF NOT EXISTS account (
			id serial PRIMARY KEY,
			first_name varchar(50),	
			last_name varchar(50),
			number serial,
			balance serial,
			created_at timestamp
		)`
	_, err := postgresStore.db.Exec(creationQuery)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(100000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(100000000)),
		CreatedAt: time.Now().UTC(),
	}
}

func (postgresStore *PostgresStore) CreateAccount(account *Account) error {
	query := `
	INSERT INTO account
	(first_name, last_name, number, balance, created_at)
	VALUES ($1, $2, $3, $4, $5)`
	resp, err := postgresStore.db.Exec(query,
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt)

	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp,
	)
	return nil
}

func (postgresStore *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (postgresStore *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (postgresStore *PostgresStore) GetAccountByIDs(int) (*Account, error) {
	return nil, nil
}
func (postgresStore *PostgresStore) GetAllAccounts() ([]*Account, error) {
	rows, err := postgresStore.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
