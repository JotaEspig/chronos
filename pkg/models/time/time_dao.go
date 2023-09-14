package time

import "database/sql"

var (
	createTimeQuery = `INSERT INTO "time"("start", "end", "repeat", "employee_id")
                       VALUES (?, ?, ?, ?);`
	findTimeByIDQuery = `SELECT "id", "start", "end", "repeat", "employee_id" FROM "time"
                         WHERE "id" = ?;`
	updateTimeQuery = `UPDATE "time" SET "start" = ?, "end" = ?, "repeat" = ?, "employee_id" = ?
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

// UpdateTime updates a time in the database
func UpdateTime(tx *sql.Tx, t *Time) error {
	stmt, err := tx.Prepare(updateTimeQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(t.Start, t.End, t.Repeat, t.EmployeeID, t.ID)
	return err
}

// DeleteTimeById deletes a user in the database by its ID
func DeleteTimeById(tx *sql.Tx, id uint) error {
	stmt, err := tx.Prepare(deleteTimeByIDQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
