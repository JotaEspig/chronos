package time

import (
	"database/sql"
	"fmt"
)

var byPage uint = 10

var (
	createTimeQuery = `INSERT INTO "time"("start", "end", "repeat", "employee_id")
                       VALUES (?, ?, ?, ?);`
	findTimeByIDQuery       = `SELECT * FROM "time" WHERE "id" = ?;`
	getNextTimesByDateQuery = fmt.Sprintf(` SELECT id FROM time WHERE "start" >= ? LIMIT %d OFFSET ?`, byPage)
	updateTimeQuery         = `UPDATE "time" SET "start" = ?, "end" = ?, "repeat" = ?, "employee_id" = ?
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

func GetNextTimesByDate(tx *sql.Tx, date string, page uint) ([]*Time, error) {
	times := make([]*Time, byPage)
	rows, err := tx.Query(getNextTimesByDateQuery, date, page*byPage)
	if err != nil {
		return []*Time{}, err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		times[i] = &Time{}
		err = rows.Scan(&times[i].ID)
		if err != nil {
			return []*Time{}, err
		}

		i++
	}

	return times, nil
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
