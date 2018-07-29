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

// DatabaseWrite interface.
type DatabaseWrite interface {
	Insert(query string, args ...interface{}) (int64, error)
	Close()
}

// DatabaseRead interface.
type DatabaseRead interface {
	Select(query string, args ...interface{}) error
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

// Select makes a SELEC.
func (db *Database) Select(query string, args ...interface{}) error {
	rows, qErr := db.db.Query(query)
	if qErr != nil {
		return qErr
	}

	defer rows.Close()
	for rows.Next() {
		rErr := rows.Scan(args...)
		if rErr != nil {
			return rErr
		}
	}
	return nil
}

// Insert store a record on the database.
func (db *Database) Insert(query string, args ...interface{}) (int64, error) {
	stmt, prepError := db.db.Prepare(query)
	if prepError != nil {
		log.Fatal(prepError)
	}

	response, respError := stmt.Exec(args...)
	if respError != nil {
		log.Fatal(respError)
	}

	return response.LastInsertId()
}

// Close ends the database connection.
func (db *Database) Close() {
	db.Close()
}
