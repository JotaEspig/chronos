// Package mtime provides support for operations with the model Time.
package time

type Time struct {
	ID         uint
	Start, End string
	Repeat     uint8
	EmployeeID uint
}
