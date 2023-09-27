package tests

import (
	"chronos/config"
	"chronos/pkg/models/time"
	"chronos/tests/timetest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeIsValid(t *testing.T) {
	timetest.TryValidTimeIsValid(t)
	timetest.TryInvalidTimeIsValid(t)
}

func TestCreateTime(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	_, err = tx.Exec(`INSERT INTO "user" VALUES (1, 'test');`)
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "employee" VALUES (1, 0, 1);`)
	assert.Nil(t, err)

	newTime := &time.Time{
		Start:      "2022-01-01",
		End:        "2022-01-01",
		Repeat:     10,
		EmployeeID: 1,
	}
	err = time.CreateTime(tx, newTime)
	assert.Nil(t, err)

	// Check that the time entry has been successfully created and has a non-zero ID.
	assert.NotEqual(t, uint(0), newTime.ID)

	// Fetch the time entry from the database by ID and check if it matches the created entry.
	fetchedT := time.Time{}
	err = tx.QueryRow(`SELECT * FROM "time"`).
		Scan(&fetchedT.ID, &fetchedT.Start, &fetchedT.End, &fetchedT.Repeat, &fetchedT.EmployeeID)
	assert.Nil(t, err)
	assert.Equal(t, newTime.ID, fetchedT.ID)
	assert.Equal(t, newTime.Start, fetchedT.Start)
	assert.Equal(t, newTime.End, fetchedT.End)
	assert.Equal(t, newTime.Repeat, fetchedT.Repeat)
	assert.Equal(t, newTime.EmployeeID, fetchedT.EmployeeID)
}

func TestFindTimeByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")
	assert.Nil(t, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "time" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Nil(t, err)

	// Try to fetch the time
	fetchedT, err := time.FindTimeByID(tx, 1)
	assert.NotNil(t, fetchedT)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), fetchedT.ID)
	assert.Equal(t, "2002-01-01", fetchedT.Start)
	assert.Equal(t, "2002-02-01", fetchedT.End)
	assert.Equal(t, uint8(1), fetchedT.Repeat)
	assert.Equal(t, uint(1), fetchedT.EmployeeID)
}

func TestUpdateTime(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")
	assert.Nil(t, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Nil(t, err)
	_, err = tx.Exec("INSERT INTO \"time\" VALUES (1, \"2022\", \"2021\", 1, 1);")
	assert.Nil(t, err)

	// Fetch the time that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"time\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to update the time
	newT := &time.Time{
		ID:         id,
		Start:      "2022",
		End:        "2022",
		Repeat:     10,
		EmployeeID: 1,
	}
	err = time.UpdateTime(tx, newT)
	assert.Nil(t, err)

	// Check if the time type has changed
	var end string
	var repeat uint8
	tx.QueryRow("SELECT \"end\", \"repeat\" FROM \"time\"").Scan(&end, &repeat)
	assert.Equal(t, newT.End, end)
	assert.Equal(t, newT.Repeat, repeat)
}

func TestDeleteTimeByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test')")
	assert.Equal(t, nil, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Equal(t, nil, err)
	_, err = tx.Exec("INSERT INTO \"time\" VALUES (1, \"2022\", \"2021\", 1, 1);")
	assert.Equal(t, nil, err)

	// Fetch the time that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"time\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to delete the time
	err = time.DeleteTimeByID(tx, id)
	assert.Equal(t, nil, err)

	// Check if the time still exists (it should not)
	id = 0
	tx.QueryRow("SELECT \"id\" FROM \"time\";").Scan(&id)
	assert.Equal(t, uint(0), id)
}
