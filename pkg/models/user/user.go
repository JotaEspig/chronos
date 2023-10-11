// package user provides support for operations with User model
package user

import (
	"chronos/pkg/types"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// InitPassword generates a bcrypt hash from password and set the password using it
func (u *User) InitPassword() {
	hashedPasswd, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPasswd)
}

// Validate validates the username and the hashed password
func (u *User) Validate(username, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return u.Username == username && err == nil
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
