package tests

import (
	"chronos/config"
	"chronos/tests/employeetest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeIsValid(t *testing.T) {
	employeetest.TryValidEmployeeIsValid(t)
	employeetest.TryInvalidEmployeeIsValid(t)
}

func TestCreateEmployee(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)
	employeetest.TryCreateValidEmployee(t, tx)
	cleanDB(tx)
	employeetest.TryCreateInvalidEmployee(t, tx)
}

func TestFindEmployeeByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)
	employeetest.TryFindValidEmployee(t, tx)
	cleanDB(tx)
	employeetest.TryFindInvalidEmployee(t, tx)
}

func TestUpdateEmployee(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)
	employeetest.TryUpdateValidEmployee(t, tx)
	cleanDB(tx)
	employeetest.TryUpdateInvalidEmployee(t, tx)
}

func TestDeleteEmployeeByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)
	employeetest.TryDeleteValidEmployee(t, tx)
}
