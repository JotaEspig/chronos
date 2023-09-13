// package user provides support for operations with User model
package user

import "database/sql"

var (
	createUserQuery         = `INSERT INTO "user"("username") VALUES (?);`
	findUserByIdQuery       = `SELECT "id", "username" FROM "user" WHERE "id" = ?;`
	findUserByUsernameQuery = `SELECT "id", "username" FROM "user"
                               WHERE "username" = ?;`
)

// CreateUser creates a user in the database
func CreateUser(tx *sql.Tx, user *User) error {
	stmt, err := tx.Prepare(createUserQuery)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.Username)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = uint(id)
	return nil
}

// FindUserByID retrieves a user from the database by its ID
func FindUserByID(tx *sql.Tx, id uint) (*User, error) {
	u := &User{}
	err := tx.QueryRow(findUserByIdQuery, id).Scan(&u.ID, &u.Username)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// FindUserByUsername retrieves a user from the database by its username
func FindUserByUsername(tx *sql.Tx, username string) (*User, error) {
	u := &User{}
	err := tx.QueryRow(findUserByUsernameQuery, username).Scan(&u.ID, &u.Username)
	if err != nil {
		return nil, err
	}

	return u, nil
}
