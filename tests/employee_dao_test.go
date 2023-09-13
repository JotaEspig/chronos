package tests

import (
	"chronos/config"
	"chronos/pkg/models/employee"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test')")

	// Create new employee
	newEmployee := &employee.Employee{
		Type:   0,
		UserID: 1,
	}

	// Insert new employee to database
	err = employee.CreateEmployee(tx, newEmployee)
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

func TestFindEmployeeByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test')")

	// Insert an employee in the database
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Equal(t, nil, err)

	// Try to fetch the employee
	e, err := employee.FindEmployeeByID(tx, 1)
	assert.NotEqual(t, nil, e)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint(1), e.ID)
	assert.Equal(t, uint8(0), e.Type)
	assert.Equal(t, uint(1), e.UserID)
}

func TestFindUserByUserID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test')")

	// Insert an employee in the database
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Equal(t, nil, err)

	// Try to fetch the employee
	e, err := employee.FindEmployeeByUserID(tx, 1)
	assert.NotEqual(t, nil, e)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint(1), e.ID)
	assert.Equal(t, uint8(0), e.Type)
	assert.Equal(t, uint(1), e.UserID)
}

func TestDeleteEmployeeByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test')")

	// Insert an employee in the database
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Equal(t, nil, err)

	// Fetch the employee that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"employee\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to delete the employee
	err = employee.DeleteEmployeeByID(tx, id)
	assert.Equal(t, nil, err)

	// Check if the employee still exists (it should not)
	id = 0
	tx.QueryRow("SELECT \"id\" FROM \"employee\";").Scan(&id)
	assert.Equal(t, uint(0), id)
}