package db

import (
	"database/sql"
	"log"
)

func NewPostgresConnection(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}


	return db
}