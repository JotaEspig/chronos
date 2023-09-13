package tests

import (
	"chronos/pkg/common"
	"database/sql"
)

// cleanDB cleans the entire DB. Be cautious when using this function
func cleanDB(tx *sql.Tx) {
	cleanQuery := common.ReadFile("./db/sql-files/clean.sql")
	tx.Exec(cleanQuery)
}
