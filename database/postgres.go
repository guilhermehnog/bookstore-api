package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "goapp"
	password = "goapp"
	host     = "127.0.0.1"
	port     = 5432
	dbname   = "library"
)

func OpenConnection() (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		panic(err)
	}
	return conn, nil
}
