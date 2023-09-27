package employee

import "database/sql"

var (
	createEmployeeQuery       = `INSERT INTO "employee"("type", "user_id") VALUES (?, ?);`
	findEmployeeByIDQuery     = `SELECT * FROM "employee" WHERE "id" = ?;`
	findEmployeeByUserIDQuery = `SELECT * FROM "employee" WHERE "user_id" = ?;`
	updateEmployeeQuery       = `UPDATE "employee" SET "type" = ?, "user_id" = ? WHERE "id" = ?;`
	deleteEmployeeByIDQuery   = `DELETE FROM "employee" WHERE "id" = ?`
)

// CreateEmployee creates an employee in the database
func CreateEmployee(tx *sql.Tx, employee *Employee) error {
	stmt, err := tx.Prepare(createEmployeeQuery)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(employee.Type, employee.UserID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	employee.ID = uint(id)
	return nil
}

// FindEmployeeByID retrieves an employee from the database by its ID
func FindEmployeeByID(tx *sql.Tx, id uint) (*Employee, error) {
	e := &Employee{}
	err := tx.QueryRow(findEmployeeByIDQuery, id).Scan(&e.ID, &e.Type, &e.UserID)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// FindEmployeeByUserID retrieves an employee from the database by its UserID
func FindEmployeeByUserID(tx *sql.Tx, id uint) (*Employee, error) {
	e := &Employee{}
	err := tx.QueryRow(findEmployeeByUserIDQuery, id).Scan(&e.ID, &e.Type, &e.UserID)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// UpdateEmployee updates an employee in the database from a given Employee
func UpdateEmployee(tx *sql.Tx, employee *Employee) error {
	stmt, err := tx.Prepare(updateEmployeeQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(employee.Type, employee.UserID, employee.ID)
	return err
}

// DeleteEmployeeByID deletes an employee in the database by its ID
func DeleteEmployeeByID(tx *sql.Tx, id uint) error {
	stmt, err := tx.Prepare(deleteEmployeeByIDQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
