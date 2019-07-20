package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

//Connection ...
type Connection struct {
	*sql.DB
}

// New ai
func New(dbTtype string, dbURL string) (*Connection, error) {

	db, err := sql.Open(dbTtype, dbURL)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxIdleConns(100)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Initiating database connection : ", dbTtype, " ", dbURL)
	return &Connection{db}, nil
}
