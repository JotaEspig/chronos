package tests

import (
	"chronos/config"
	"chronos/pkg/models/scheduling"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateScheduling(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	_, err = tx.Exec(`INSERT INTO "user" VALUES (1, 'test');`)
	assert.Equal(t, nil, err)
	_, err = tx.Exec(`INSERT INTO "employee" VALUES (1, 0, 1);`)
	assert.Equal(t, nil, err)
	_, err = tx.Exec(`INSERT INTO "time" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Equal(t, nil, err)

	newScheduling := &scheduling.Scheduling{
		Start:  "2022-01-01",
		End:    "2022-01-01",
		UserID: 1,
		TimeID: 1,
	}
	err = scheduling.CreateScheduling(tx, newScheduling)
	assert.Equal(t, nil, err)

	// Check that the scheduling entry has been successfully created and has a non-zero ID.
	assert.NotEqual(t, uint(0), newScheduling.ID)

	// Fetch the scheduling entry from the database by ID and check if it matches the created entry.
	sched := scheduling.Scheduling{}
	err = tx.QueryRow(`SELECT * FROM "scheduling"`).
		Scan(&sched.ID, &sched.Start, &sched.End, &sched.UserID, &sched.TimeID)
	assert.Equal(t, nil, err)
	assert.Equal(t, newScheduling.ID, sched.ID)
	assert.Equal(t, newScheduling.Start, sched.Start)
	assert.Equal(t, newScheduling.End, sched.End)
	assert.Equal(t, newScheduling.UserID, sched.UserID)
	assert.Equal(t, newScheduling.TimeID, sched.TimeID)
}

func TestFindSchedulingByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")
	assert.Equal(t, nil, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Equal(t, nil, err)
	_, err = tx.Exec(`INSERT INTO "time" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Equal(t, nil, err)
	_, err = tx.Exec(`INSERT INTO "scheduling" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Equal(t, nil, err)

	// Try to fetch the scheduling
	s, err := scheduling.FindSchedulingByID(tx, 1)
	assert.NotEqual(t, nil, s)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint(1), s.ID)
	assert.Equal(t, "2002-01-01", s.Start)
	assert.Equal(t, "2002-02-01", s.End)
	assert.Equal(t, uint(1), s.UserID)
	assert.Equal(t, uint(1), s.TimeID)
}

func TestUpdateScheduling(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")
	assert.Equal(t, nil, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Equal(t, nil, err)
	_, err = tx.Exec("INSERT INTO \"time\" VALUES (1, \"2022\", \"2021\", 1, 1);")
	assert.Equal(t, nil, err)
	_, err = tx.Exec(`INSERT INTO "scheduling" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Equal(t, nil, err)

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
	assert.Equal(t, nil, err)

	// Check if the time type has changed
	var start string
	var end string
	tx.QueryRow("SELECT \"start\", \"end\" FROM \"scheduling\"").Scan(&start, &end)
	assert.Equal(t, s.Start, start)
	assert.Equal(t, s.End, end)
}

func TestDeleteSchedulingByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test')")
	assert.Equal(t, nil, err)
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Equal(t, nil, err)
	_, err = tx.Exec("INSERT INTO \"time\" VALUES (1, \"2022\", \"2021\", 1, 1);")
	assert.Equal(t, nil, err)
	_, err = tx.Exec(`INSERT INTO "scheduling" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Equal(t, nil, err)

	// Fetch the scheduling that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"scheduling\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to delete the scheduling
	err = scheduling.DeleteSchedulingByID(tx, id)
	assert.Equal(t, nil, err)

	// Check if the scheduling still exists (it should not)
	id = 0
	tx.QueryRow("SELECT \"id\" FROM \"scheduling\";").Scan(&id)
	assert.Equal(t, uint(0), id)
}
