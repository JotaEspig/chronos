package config

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./db.db")
	if err != nil {
		panic(err)
	}
}

func init() {
	initDB()
}
