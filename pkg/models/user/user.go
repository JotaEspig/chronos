// package user provides support for operations with User model
package user

import (
	"chronos/pkg/types"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

const (
	TypeStudent  uint8 = 0
	TypeEmployee uint8 = 1
	TypeAdmin    uint8 = 2
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Type     uint8  `json:"type"`
	Password string `json:"password"`
}

// InitPassword generates a bcrypt hash from password and set the password using it
func (u *User) InitPassword() error {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPasswd)
	return err
}

// ValidateLogin validates the username and the hashed password
func (u *User) ValidateLogin(username, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return u.Username == username && err == nil
}

func (u *User) IsValid() bool {
	return u.Username != "" && u.Password != "" &&
		(u.Type >= TypeStudent && u.Type <= TypeAdmin)
}

func (u *User) Sanitize(policy *bluemonday.Policy) {
	u.Username = policy.Sanitize(u.Username)
}

func (u *User) ToMap() types.JsonMap {
	m := make(types.JsonMap)
	m["id"] = u.ID
	m["username"] = u.Username
	m["type"] = u.Type
	return m
}
