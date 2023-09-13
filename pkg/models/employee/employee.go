// package employee provides support for operations with Employee model
package employee

type Employee struct {
	ID     uint
	Type   uint8
	UserID uint
}

func (e *Employee) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = e.ID
	m["type"] = e.Type
	m["user_id"] = e.UserID
	return m
}
