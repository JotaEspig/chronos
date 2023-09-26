// package time provides support for operations with the model Time.
package time

import (
	"chronos/pkg/types"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

type RepeatEnum uint8

// Based on UNIX permission
// These enums are simple. Maybe it should be refined
const (
	Monday    RepeatEnum = 0b00000001
	Tuesday   RepeatEnum = 0b00000010
	Wednesday RepeatEnum = 0b00000100
	Thursday  RepeatEnum = 0b00001000
	Friday    RepeatEnum = 0b00010000
	Daily     RepeatEnum = 0b00100000
	Weekly    RepeatEnum = 0b01000000
)

type Time struct {
	ID         uint   `json:"id"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Repeat     uint8  `json:"repeat"`
	EmployeeID uint   `json:"employee_id"`
}

func (t *Time) IsValid() bool {
	start, err := time.Parse(time.DateTime, t.Start)
	validations := err == nil
	end, err := time.Parse(time.DateTime, t.End)
	validations = validations && end.After(start)
	validations = validations && err == nil

	return validations && t.EmployeeID != 0
}

func (t *Time) Sanitize(policy *bluemonday.Policy) {
	t.Start = policy.Sanitize(t.Start)
	t.End = policy.Sanitize(t.End)
}

func (t *Time) ToMap() types.JsonMap {
	m := make(types.JsonMap)
	m["id"] = t.ID
	m["start"] = t.Start
	m["end"] = t.End
	m["repeat"] = t.Repeat
	m["employee_id"] = t.EmployeeID
	return m
}
