package timetest

import (
	"chronos/pkg/models/time"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TryGetTimesByDate(t *testing.T, tx *sql.Tx) {
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 1, 'test1');")
	tx.Exec("INSERT INTO \"employee\" VALUES (1, 1, 1);")
	tx.Exec("INSERT INTO \"time\" VALUES (1, '2022-01-03', '2022-12-01', 64, 1);")
	tx.Exec("INSERT INTO \"time\" VALUES (2, '2022-01-10', '2022-12-01', 64, 1);")
	tx.Exec("INSERT INTO \"time\" VALUES (3, '2022-01-17', '2022-12-01', 64, 1);")
	// should not
	tx.Exec("INSERT INTO \"time\" VALUES (4, '2022-01-12', '2022-12-01', 64, 1);")
	// should not
	tx.Exec("INSERT INTO \"time\" VALUES (5, '2022-01-20', '2022-12-01', 64, 1);")
	tx.Exec("INSERT INTO \"time\" VALUES (6, '2022-01-24', '2022-12-01', 32, 1);")
	// here is 1 because 1 means monday in time.go and 2022-01-03 is monday
	tx.Exec("INSERT INTO \"time\" VALUES (7, '2022-01-10', '2022-12-01', 1, 1);")
	// should not
	tx.Exec("INSERT INTO \"time\" VALUES (8, '2022-01-10', '2022-12-01', 2, 1);")

	times, err := time.GetTimesByDate(tx, "2022-01-03", 0)
	assert.Nil(t, err)
	assert.Equal(t, 5, len(times))
	for _, elem := range times {
		fmt.Println(elem)
	}
}
