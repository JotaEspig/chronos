package employeetest

import (
	"chronos/pkg/models/employee"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TryCreateValidEmployee(t *testing.T, tx *sql.Tx) {
	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")

	// Create new employee
	newEmployee := &employee.Employee{
		Type:   0,
		UserID: 1,
	}

	// Insert new employee to database
	err := employee.CreateEmployee(tx, newEmployee)
	assert.Nil(t, err)

	// Check that the employee has been successfully created and has a non-zero ID.
	assert.NotEqual(t, uint(0), newEmployee.ID)

	// Fetch the employee from the database by ID and check if it matches the created employee.
	var id uint
	var _type uint8
	var userID uint

	// Check if it was created
	tx.QueryRow("SELECT \"id\", \"type\", \"user_id\" FROM \"employee\";").Scan(&id, &_type, &userID)
	assert.Nil(t, err)
	assert.Equal(t, newEmployee.ID, id)
	assert.Equal(t, newEmployee.Type, _type)
	assert.Equal(t, newEmployee.UserID, userID)
}

func TryCreateInvalidEmployee(t *testing.T, tx *sql.Tx) {
	// Create new employee
	newEmployee := &employee.Employee{
		Type:   1,
		UserID: 1,
	}

	// Try to insert new employee to database
	err := employee.CreateEmployee(tx, newEmployee)
	assert.NotNil(t, err)

	// Check that the employee has been successfully created and has a non-zero ID.
	assert.Equal(t, uint(0), newEmployee.ID)

	// Fetch the employee from the database by ID and check if it matches the created employee.
	var id uint
	var _type uint8
	var userID uint

	// Check if it was created
	tx.QueryRow("SELECT \"id\", \"type\", \"user_id\" FROM \"employee\";").Scan(&id, &_type, &userID)
	assert.NotNil(t, err)
	assert.NotEqual(t, newEmployee.Type, _type)
	assert.NotEqual(t, newEmployee.UserID, userID)
}

func TryFindValidEmployee(t *testing.T, tx *sql.Tx) {
	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")

	// Insert an employee in the database
	_, err := tx.Exec("INSERT INTO \"employee\" VALUES (1, 2, 1);")
	assert.Nil(t, err)

	// Try to fetch the employee
	e, err := employee.FindEmployeeByID(tx, 1)
	assert.NotNil(t, e)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), e.ID)
	assert.Equal(t, uint8(2), e.Type)
	assert.Equal(t, uint(1), e.UserID)
}

func TryFindInvalidEmployee(t *testing.T, tx *sql.Tx) {
	// Try to fetch the employee (it shouldn't work)
	e, err := employee.FindEmployeeByID(tx, 1)
	assert.NotNil(t, err)
	assert.Nil(t, e)
}

func TryUpdateValidEmployee(t *testing.T, tx *sql.Tx) {
	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")

	// Insert an employee in the database
	_, err := tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Nil(t, err)

	// Fetch the employee that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"employee\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to update the employee
	e := &employee.Employee{
		ID:     id,
		Type:   1,
		UserID: 1,
	}
	err = employee.UpdateEmployee(tx, e)
	assert.Nil(t, err)

	// Check if the employee type has changed
	var _type uint8
	tx.QueryRow("SELECT \"type\" FROM \"employee\"").Scan(&_type)
	assert.Equal(t, e.Type, _type)
}

func TryUpdateInvalidEmployee(t *testing.T, tx *sql.Tx) {
	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")

	// Insert an employee in the database
	_, err := tx.Exec("INSERT INTO \"employee\" VALUES (1, 2, 1);")
	assert.Nil(t, err)

	// Fetch the employee that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"employee\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to update the employee
	e := &employee.Employee{
		ID:     id,
		Type:   1,
		UserID: 0,
	}
	err = employee.UpdateEmployee(tx, e)
	assert.NotNil(t, err)

	// Check if the employee type has changed
	var _type uint8
	tx.QueryRow("SELECT \"type\" FROM \"employee\"").Scan(&_type)
	assert.NotEqual(t, e.Type, _type)
}

func TryDeleteValidEmployee(t *testing.T, tx *sql.Tx) {
	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")

	// Insert an employee in the database
	_, err := tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Nil(t, err)

	// Fetch the employee that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"employee\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to delete the employee
	err = employee.DeleteEmployeeByID(tx, id)
	assert.Nil(t, err)

	// Check if the employee still exists (it should not)
	id = 0
	tx.QueryRow("SELECT \"id\" FROM \"employee\";").Scan(&id)
	assert.Equal(t, uint(0), id)
}
