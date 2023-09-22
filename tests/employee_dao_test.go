package tests

import (
	"chronos/config"
	"chronos/pkg/models/employee"
	"chronos/tests/employee_test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)
	employee_test.TryCreateValidEmployee(t, tx)
	cleanDB(tx)
	employee_test.TryCreateInvalidEmployee(t, tx)
}

func TestFindEmployeeByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")

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
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")

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

func TestUpdateEmployee(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a placeholder user
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")

	// Insert an employee in the database
	_, err = tx.Exec("INSERT INTO \"employee\" VALUES (1, 0, 1);")
	assert.Equal(t, nil, err)

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
	assert.Equal(t, nil, err)

	// Check if the employee type has changed
	var _type uint8
	tx.QueryRow("SELECT \"type\" FROM \"employee\"").Scan(&_type)
	assert.Equal(t, e.Type, _type)
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
