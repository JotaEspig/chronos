package employee_test

import (
	"chronos/pkg/models/employee"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TryCreateValidEmployee(t *testing.T, tx *sql.Tx) {
	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")

	// Create new employee
	newEmployee := &employee.Employee{
		Type:   0,
		UserID: 1,
	}

	// Insert new employee to database
	err := employee.CreateEmployee(tx, newEmployee)
	assert.Equal(t, nil, err)

	// Check that the employee has been successfully created and has a non-zero ID.
	assert.NotEqual(t, uint(0), newEmployee.ID)

	// Fetch the employee from the database by ID and check if it matches the created employee.
	var id uint
	var _type uint8
	var userID uint

	// Check if it was created
	tx.QueryRow("SELECT \"id\", \"type\", \"user_id\" FROM \"employee\";").Scan(&id, &_type, &userID)
	assert.Equal(t, nil, err)
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
	assert.NotEqual(t, nil, err)

	// Check that the employee has been successfully created and has a non-zero ID.
	assert.Equal(t, uint(0), newEmployee.ID)

	// Fetch the employee from the database by ID and check if it matches the created employee.
	var id uint
	var _type uint8
	var userID uint

	// Check if it was created
	tx.QueryRow("SELECT \"id\", \"type\", \"user_id\" FROM \"employee\";").Scan(&id, &_type, &userID)
	assert.NotEqual(t, nil, err)
	assert.NotEqual(t, newEmployee.Type, _type)
	assert.NotEqual(t, newEmployee.UserID, userID)
}
