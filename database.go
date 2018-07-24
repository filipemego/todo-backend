package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Database struct.
type Database struct {
	db *sql.DB
}

// DatabaseWrite interface
type DatabaseWrite interface {
	Close()
}

// DatabaseRead interface
type DatabaseRead interface {
	Close()
}

// NewDatabase constructor.
func NewDatabase(databaseURI string) *Database {
	db, openErr := sql.Open("mysql", databaseURI)
	if openErr != nil {
		log.Fatal(openErr)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return &Database{db: db}
}

// Close ends the database connection
func (db *Database) Close() {
	db.Close()
}
