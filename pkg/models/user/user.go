// package user provides support for operations with User model
package user

import (
	"chronos/pkg/types"

	"github.com/microcosm-cc/bluemonday"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func (u *User) IsValid() bool {
	return u.Username != ""
}

func (u *User) Sanitize(policy *bluemonday.Policy) {
	u.Username = policy.Sanitize(u.Username)
}

func (u *User) ToMap() types.JsonMap {
	m := make(types.JsonMap)
	m["id"] = u.ID
	m["username"] = u.Username
	return m
}
