package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
}
