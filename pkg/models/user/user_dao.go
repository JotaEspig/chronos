// package user provides support for operations with User model
package user

import "chronos/config"

var (
	db                = config.DB
	findUserByIdQuery = `SELECT "id", "username" FROM "user" WHERE "id" = ?;`
)

func FindUserByID(id uint) *User {
	u := &User{}
	err := db.QueryRow(findUserByIdQuery, id).Scan(&u.ID, &u.Username)
	if err != nil {
		return nil
	}

	return u
}
