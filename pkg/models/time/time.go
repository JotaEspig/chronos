// Package mtime provides support for operations with the model Time.
package time

import "chronos/pkg/types"

type Time struct {
	ID         uint   `json:"id"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Repeat     uint8  `json:"repeat"`
	EmployeeID uint   `json:"employee_id"`
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
