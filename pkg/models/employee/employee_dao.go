// package employee provides support for operations with Employee model
package employee

import "database/sql"

var (
	createEmployeeQuery       = `INSERT INTO "employee"("type", "user_id") VALUES (?, ?);`
	findEmployeeByIDQuery     = `SELECT "id", "type", "user_id" FROM "employee" WHERE "id" = ?;`
	findEmployeeByUserIDQuery = `SELECT "id", "type", "user_id" FROM "employee" WHERE "user_id" = ?;`
	deleteEmployeeByIDQuery   = `DELETE FROM "employee" WHERE "id" = ?`
)

// CreateEmployee creates a user in the database
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
	u := &Employee{}
	err := tx.QueryRow(findEmployeeByIDQuery, id).Scan(&u.ID, &u.Type, &u.UserID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// FindEmployeeByUserID retrieves an employee from the database by its UserID
func FindEmployeeByUserID(tx *sql.Tx, id uint) (*Employee, error) {
	u := &Employee{}
	err := tx.QueryRow(findEmployeeByUserIDQuery, id).Scan(&u.ID, &u.Type, &u.UserID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// DeleteEmployeeByID deletes an employee from the database by its ID
func DeleteEmployeeByID(tx *sql.Tx, id uint) error {
	stmt, err := tx.Prepare(deleteEmployeeByIDQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	return err
}
