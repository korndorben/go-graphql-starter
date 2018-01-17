package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func OpenDB(path string) (*sqlx.DB, error) {
	log.Println("Database is inilializing... ")
	db, err := sqlx.Open("sqlite3", path)
	log.Println("Database is inilialized ")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
