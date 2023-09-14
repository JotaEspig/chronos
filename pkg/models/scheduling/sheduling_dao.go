package scheduling

import "database/sql"

var (
	createSchedulingQuery = `INSERT INTO "scheduling"("start", "end", "user_id", "time_id")
                       VALUES (?, ?, ?, ?);`
	findSchedulingByIDQuery = `SELECT "id", "start", "end", "user_id", "time_id" FROM "scheduling"
                         WHERE "id" = ?;`
	updateSchedulingQuery = `UPDATE "scheduling" SET "start" = ?, "end" = ?, "user_id" = ?, "time_id" = ?
                       WHERE "id" = ?;`
	deleteSchedulingByIDQuery = `DELETE FROM "scheduling" WHERE "id" = ?;`
)

// CreateScheduling creates a scheduling in the database
func CreateScheduling(tx *sql.Tx, sched *Scheduling) error {
	stmt, err := tx.Prepare(createSchedulingQuery)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(sched.Start, sched.End, sched.UserID, sched.TimeID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	sched.ID = uint(id)
	return nil
}

// FindSchedulingByID retrieves a scheduling from the database by its ID
func FindSchedulingByID(tx *sql.Tx, id uint) (*Scheduling, error) {
	s := &Scheduling{}
	err := tx.QueryRow(findSchedulingByIDQuery, id).
		Scan(&s.ID, &s.Start, &s.End, &s.UserID, &s.TimeID)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateScheduling updates a scheduling in the database
func UpdateScheduling(tx *sql.Tx, s *Scheduling) error {
	stmt, err := tx.Prepare(updateSchedulingQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(s.Start, s.End, s.UserID, s.TimeID, s.ID)
	return err
}

// DeleteSchedulingByID deletes a scheduling in the database by its ID
func DeleteSchedulingByID(tx *sql.Tx, id uint) error {
	stmt, err := tx.Prepare(deleteSchedulingByIDQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
