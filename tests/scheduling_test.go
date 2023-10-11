package tests

import (
	"chronos/config"
	"chronos/pkg/models/scheduling"
	"chronos/tests/schedulingtest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateScheduling(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)
	schedulingtest.TryCreateValidScheduling(t, tx)
	cleanDB(tx)
	schedulingtest.TryCreateInvalidScheduling(t, tx)
}

func TestFindSchedulingByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")
	assert.Nil(t, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "time" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "scheduling" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Nil(t, err)

	// Try to fetch the scheduling
	s, err := scheduling.FindSchedulingByID(tx, 1)
	assert.NotNil(t, s)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), s.ID)
	assert.Equal(t, "2002-01-01", s.Start)
	assert.Equal(t, "2002-02-01", s.End)
	assert.Equal(t, uint(1), s.UserID)
	assert.Equal(t, uint(1), s.TimeID)
}

func TestUpdateScheduling(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")
	assert.Nil(t, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Nil(t, err)
	_, err = tx.Exec("INSERT INTO \"time\" VALUES (1, \"2022\", \"2021\", 1, 1);")
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "scheduling" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Nil(t, err)

	// Fetch the scheduling that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"scheduling\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to update the scheduling
	s := &scheduling.Scheduling{
		ID:     id,
		Start:  "2022",
		End:    "2022",
		UserID: 1,
		TimeID: 1,
	}
	err = scheduling.UpdateScheduling(tx, s)
	assert.Nil(t, err)

	// Check if the time type has changed
	var start string
	var end string
	tx.QueryRow("SELECT \"start\", \"end\" FROM \"scheduling\"").Scan(&start, &end)
	assert.Equal(t, s.Start, start)
	assert.Equal(t, s.End, end)
}

func TestDeleteSchedulingByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")
	assert.Nil(t, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Nil(t, err)
	_, err = tx.Exec("INSERT INTO \"time\" VALUES (1, \"2022\", \"2021\", 1, 1);")
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "scheduling" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Nil(t, err)

	// Fetch the scheduling that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"scheduling\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to delete the scheduling
	err = scheduling.DeleteSchedulingByID(tx, id)
	assert.Nil(t, err)

	// Check if the scheduling still exists (it should not)
	id = 0
	tx.QueryRow("SELECT \"id\" FROM \"scheduling\";").Scan(&id)
	assert.Equal(t, uint(0), id)
}
