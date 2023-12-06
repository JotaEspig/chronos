package time

import (
	"chronos/pkg/common"
	"database/sql"
	"fmt"
)

var byPage uint = 10

var (
	createTimeQuery = `INSERT INTO "time"("start", "end", "repeat", "employee_id")
                       VALUES (?, ?, ?, ?);`
	findTimeByIDQuery   = `SELECT * FROM "time" WHERE "id" = ?;`
	getTimesByDateQuery = "" // set in GetNextTimesByDate func
	updateTimeQuery     = `UPDATE "time" SET "start" = ?, "end" = ?, "repeat" = ?, "employee_id" = ?
                       WHERE "id" = ?;`
	deleteTimeByIDQuery = `DELETE FROM "time" WHERE "id" = ?;`
)

// CreateTime creates a time in the database
func CreateTime(tx *sql.Tx, time *Time) error {
	stmt, err := tx.Prepare(createTimeQuery)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(time.Start, time.End, time.Repeat, time.EmployeeID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	time.ID = uint(id)
	return nil
}

// FindTimeByID retrieves a time from the database by its ID
func FindTimeByID(tx *sql.Tx, id uint) (*Time, error) {
	t := &Time{}
	err := tx.QueryRow(findTimeByIDQuery, id).
		Scan(&t.ID, &t.Start, &t.End, &t.Repeat, &t.EmployeeID)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func GetTimesByDate(tx *sql.Tx, date string, page uint) ([]*Time, error) {
	if getTimesByDateQuery == "" {
		getTimesByDateQuery = fmt.Sprintf(
			common.ReadFile("./db/sql-files/queries-with-params/get_times_by_date.sql"),
			byPage,
		)
	}

	times := make([]*Time, byPage)
	rows, err := tx.Query(getTimesByDateQuery, date, date, date, date, page*byPage)
	if err != nil {
		return []*Time{}, err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		times[i] = &Time{}
		err = rows.Scan(
			&times[i].ID, &times[i].Start, &times[i].End,
			&times[i].Repeat, &times[i].EmployeeID,
		)
		if err != nil {
			return []*Time{}, err
		}

		i++
	}

	return times[:i], nil
}

// UpdateTime updates a time in the database
func UpdateTime(tx *sql.Tx, t *Time) error {
	stmt, err := tx.Prepare(updateTimeQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(t.Start, t.End, t.Repeat, t.EmployeeID, t.ID)
	return err
}

// DeleteTimeByID deletes a time in the database by its ID
func DeleteTimeByID(tx *sql.Tx, id uint) error {
	stmt, err := tx.Prepare(deleteTimeByIDQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
