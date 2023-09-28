package employeetest

import (
	"chronos/pkg/models/employee"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TryValidEmployeeIsValid(t *testing.T) {
	e := employee.Employee{
		Type:   1,
		UserID: 1,
	}
	assert.Equal(t, true, e.IsValid())
}

func TryInvalidEmployeeIsValid(t *testing.T) {
	e := employee.Employee{
		Type:   1,
		UserID: 0,
	}
	assert.Equal(t, false, e.IsValid())
}
