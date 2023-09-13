// package user provides support for operations with User model
package user

import "database/sql"

var (
	findUserByIdQuery = `SELECT "id", "username" FROM "user" WHERE "id" = ?;`
)

func FindUserByID(tx *sql.Tx, id uint) *User {
	u := &User{}
	err := tx.QueryRow(findUserByIdQuery, id).Scan(&u.ID, &u.Username)
	if err != nil {
		return nil
	}

	return u
}
