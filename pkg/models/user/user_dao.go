package user

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	createUserQuery         = `INSERT INTO "user"("username") VALUES (?);`
	findUserByIDQuery       = `SELECT "id", "username" FROM "user" WHERE "id" = ?;`
	findUserByUsernameQuery = `SELECT "id", "username" FROM "user"
                               WHERE "username" = ?;`
	updateUserQuery     = `UPDATE "user" SET "username" = ? WHERE "id" = ?;`
	deleteUserByIDQuery = `DELETE FROM "user" WHERE "id" = ?;`
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
	err := tx.QueryRow(findUserByIDQuery, id).Scan(&u.ID, &u.Username)
	if err != nil {
		return nil, err
	}
	if u.ID == 0 {
		errorString := fmt.Sprintf("user_dao: user with id = %d not found", u.ID)
		return nil, errors.New(errorString)
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

// UpdateUser updates a user in the database
func UpdateUser(tx *sql.Tx, u *User) error {
	stmt, err := tx.Prepare(updateUserQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Username, u.ID)
	return err
}

// DeleteUserByID deletes a user in the database by its ID
func DeleteUserByID(tx *sql.Tx, id uint) error {
	stmt, err := tx.Prepare(deleteUserByIDQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
