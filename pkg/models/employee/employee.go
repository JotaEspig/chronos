// package employee provides support for operations with Employee model
package employee

import "chronos/pkg/types"

type Employee struct {
	ID     uint  `json:"id"`
	Type   uint8 `json:"type"`
	UserID uint  `json:"user_id"`
}

func (e *Employee) ToMap() types.JsonMap {
	m := make(types.JsonMap)
	m["id"] = e.ID
	m["type"] = e.Type
	m["user_id"] = e.UserID
	return m
}
