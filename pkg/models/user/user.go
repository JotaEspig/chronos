// package user provides support for operations with User model
package user

type User struct {
	ID       uint
	Username string
}

func (u *User) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = u.ID
	m["username"] = u.Username
	return m
}
