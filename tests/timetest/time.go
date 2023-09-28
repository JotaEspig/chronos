package timetest

import (
	"chronos/pkg/models/time"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TryValidTimeIsValid(t *testing.T) {
	_t := time.Time{
		Start:      "2020-01-01 12:30:20",
		End:        "2020-01-01 12:30:21",
		EmployeeID: 1,
	}
	assert.Equal(t, true, _t.IsValid())
	_t = time.Time{
		Start:      "2020-01-01 10:30:20",
		End:        "2020-01-01 12:30:20",
		EmployeeID: 1,
	}
	assert.Equal(t, true, _t.IsValid())
	_t = time.Time{
		Start:      "2025-01-01 10:30:20",
		End:        "2025-02-01 10:30:20",
		EmployeeID: 1,
	}
	assert.Equal(t, true, _t.IsValid())
}

func TryInvalidTimeIsValid(t *testing.T) {
	_t := time.Time{
		Start:      "2020-01-01 12:0:20",
		End:        "2021-01-01 12:30:20",
		EmployeeID: 1,
	}
	assert.Equal(t, false, _t.IsValid())
	_t = time.Time{
		Start:      "2020-01-01 1:3020",
		End:        "2020-01-01 12:3:20",
		EmployeeID: 1,
	}
	assert.Equal(t, false, _t.IsValid())
	_t = time.Time{
		Start:      "2020-01-01 11:30:20",
		End:        "2020-0-01 12:3:20",
		EmployeeID: 1,
	}
	assert.Equal(t, false, _t.IsValid())
	_t = time.Time{
		Start:      "1020-01-01 13020",
		End:        "00-0-01 12:3:20",
		EmployeeID: 1,
	}
	assert.Equal(t, false, _t.IsValid())
	_t = time.Time{
		Start:      "2020-01-01 10:30:20",
		End:        "2020-01-01 12:30:20",
		EmployeeID: 0,
	}
	assert.Equal(t, false, _t.IsValid())
}
