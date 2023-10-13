package schedulingtest

import (
	"chronos/pkg/models/scheduling"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TryCreateValidScheduling(t *testing.T, tx *sql.Tx) {
	_, err := tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 1, 'test1');")
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "employee" VALUES (1, 0, 1);`)
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "time" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Nil(t, err)

	newScheduling := &scheduling.Scheduling{
		Start:  "2022-01-01",
		End:    "2022-01-01",
		UserID: 1,
		TimeID: 1,
	}
	err = scheduling.CreateScheduling(tx, newScheduling)
	assert.Nil(t, err)

	// Check that the scheduling entry has been successfully created and has a non-zero ID.
	assert.NotEqual(t, uint(0), newScheduling.ID)

	// Fetch the scheduling entry from the database by ID and check if it matches the created entry.
	sched := scheduling.Scheduling{}
	err = tx.QueryRow(`SELECT * FROM "scheduling"`).
		Scan(&sched.ID, &sched.Start, &sched.End, &sched.UserID, &sched.TimeID)
	assert.Nil(t, err)
	assert.Equal(t, newScheduling.ID, sched.ID)
	assert.Equal(t, newScheduling.Start, sched.Start)
	assert.Equal(t, newScheduling.End, sched.End)
	assert.Equal(t, newScheduling.UserID, sched.UserID)
	assert.Equal(t, newScheduling.TimeID, sched.TimeID)
}

func TryCreateInvalidScheduling(t *testing.T, tx *sql.Tx) {
	_, err := tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 1, 'test1');")
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "employee" VALUES (1, 0, 1);`)
	assert.Nil(t, err)
	_, err = tx.Exec(`INSERT INTO "time" VALUES (1, "2002-01-01", "2002-02-01", 1, 1);`)
	assert.Nil(t, err)

	newScheduling := &scheduling.Scheduling{
		Start:  "2022-01-01",
		End:    "2022-01-01",
		UserID: 0,
		TimeID: 1,
	}
	err = scheduling.CreateScheduling(tx, newScheduling)
	assert.NotNil(t, err)

	// Check that the scheduling entry has been successfully created and has a non-zero ID.
	assert.Equal(t, uint(0), newScheduling.ID)

	// Fetch the scheduling entry from the database by ID and check if it matches the created entry.
	sched := scheduling.Scheduling{}
	err = tx.QueryRow(`SELECT * FROM "scheduling"`).
		Scan(&sched.ID, &sched.Start, &sched.End, &sched.UserID, &sched.TimeID)
	assert.NotNil(t, err)
	assert.NotEqual(t, newScheduling.Start, sched.Start)
	assert.NotEqual(t, newScheduling.End, sched.End)
	assert.Equal(t, uint(0), sched.UserID)
	assert.NotEqual(t, newScheduling.TimeID, sched.TimeID)
}
