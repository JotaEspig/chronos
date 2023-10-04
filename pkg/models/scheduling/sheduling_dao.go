package scheduling

import (
	"chronos/pkg/common"
	"database/sql"
	"fmt"
)

var byPage uint = 10

var (
	createSchedulingQuery = `INSERT INTO "scheduling"("start", "end", "user_id", "time_id")
                       VALUES (?, ?, ?, ?);`
	findSchedulingByIDQuery   = `SELECT * FROM "scheduling" WHERE "id" = ?;`
	getSchedulingsByDateQuery = ""
	updateSchedulingQuery     = `UPDATE "scheduling" SET "start" = ?, "end" = ?, "user_id" = ?, "time_id" = ?
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

func GetSchedulingsByDate(tx *sql.Tx, date string, page uint) ([]*Scheduling, error) {
	if getSchedulingsByDateQuery == "" {
		getSchedulingsByDateQuery = fmt.Sprintf(
			common.ReadFile("./db/sql-files/queries-with-params/get_schedulings_by_date.sql"),
			byPage,
		)
	}

	schedulings := make([]*Scheduling, byPage)
	rows, err := tx.Query(getSchedulingsByDateQuery, date, page*byPage)
	if err != nil {
		return []*Scheduling{}, err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		schedulings[i] = &Scheduling{}
		err = rows.Scan(
			&schedulings[i].ID, &schedulings[i].Start, &schedulings[i].End,
			&schedulings[i].UserID, &schedulings[i].TimeID,
		)
		if err != nil {
			return []*Scheduling{}, err
		}

		i++
	}

	return schedulings[:i], nil
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
