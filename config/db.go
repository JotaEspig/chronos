package config

import (
	"chronos/pkg/common"
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func createTables() {
	filePath := ""
	if common.IsTestRun() {
		filePath += "../"
	}
	filePath += "db/sql-files/create_tables.sql"

	query := common.ReadFile(filePath)
	DB.Exec(query)
}

func InitDB() {
	var err error
	dbString := os.Getenv("CHRONOS_DB_STRING")
	if dbString == "" {
		panic(errors.New("CHRONOS_DB_STRING is not set"))
	}

	DB, err = sql.Open("sqlite3", dbString)
	if err != nil {
		panic(err)
	}

	createTables()
}
