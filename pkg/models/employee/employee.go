// package employee provides support for operations with Employee model
package employee

import "chronos/pkg/types"

type TypeEnum uint8

// Based on UNIX permission
// These are placeholders, not the actual types.
// TODO: make it correct
const (
	TypeProfessor    = 0b00000001
	TypePsychologist = 0b00000010
)

type Employee struct {
	ID     uint  `json:"id"`
	Type   uint8 `json:"type"`
	UserID uint  `json:"user_id"`
}

func (e *Employee) IsValid() bool {
	return e.UserID != 0
}

func (e *Employee) ToMap() types.JsonMap {
	m := make(types.JsonMap)
	m["id"] = e.ID
	m["type"] = e.Type
	m["user_id"] = e.UserID
	return m
}
