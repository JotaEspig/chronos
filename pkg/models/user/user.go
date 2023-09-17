// package user provides support for operations with User model
package user

import "chronos/pkg/types"

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func (u *User) IsValid() bool {
	return u.Username != ""
}

func (u *User) ToMap() types.JsonMap {
	m := make(types.JsonMap)
	m["id"] = u.ID
	m["username"] = u.Username
	return m
}
