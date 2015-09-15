package models

import (
	"database/sql"
	"os"

	_ "github.com/goodeggs/ccron/ccron-api/Godeps/_workspace/src/github.com/lib/pq"
)

var Db *sql.DB

func Connect() error {
	db_url := os.Getenv("POSTGRES_URL")
	if db_url == "" {
		db_url = "postgres://ccron:ccron@localhost/ccron_development?sslmode=disable"
	}

	db, err := sql.Open("postgres", db_url)

	if err != nil {
		return err
	}

	Db = db

	Db.Exec(`CREATE TABLE jobs (
		id SERIAL,
		app TEXT,
		schedule TEXT,
		command TEXT,
		PRIMARY KEY (id)
	);`)

	return nil
}

func Disconnect() error {
	return Db.Close()
}
