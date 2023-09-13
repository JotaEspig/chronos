// Package mtime provides support for operations with the model Time.
package time

type Time struct {
	ID         uint
	Start, End string
	Repeat     uint8
	EmployeeID uint
}

func (t *Time) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = t.ID
	m["start"] = t.Start
	m["end"] = t.End
	m["repeat"] = t.Repeat
	m["employee_id"] = t.EmployeeID
	return m
}
