package tests

import (
	"chronos/config"
	"chronos/pkg/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	newUser := &user.User{
		Username: "newuser",
		Password: "test",
	}

	err = user.CreateUser(tx, newUser)
	assert.Nil(t, err)

	// Check that the user has been successfully created and has a non-zero ID.
	assert.NotEqual(t, uint(0), newUser.ID)

	// Fetch the user from the database by ID and check if it matches the created user.
	var id uint
	var username string
	var password string

	// Check if it was created
	tx.QueryRow("SELECT * FROM \"user\";").Scan(&id, &username, &password)
	assert.Nil(t, err)
	assert.Equal(t, newUser.ID, id)
	assert.Equal(t, newUser.Username, username)
	assert.Equal(t, newUser.Password, password)
}

func TestFindUserByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a user in the database
	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")
	assert.Nil(t, err)

	// Try to fetch the user
	u, err := user.FindUserByID(tx, 1)
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), u.ID)
	assert.Equal(t, "test", u.Username)
	assert.Equal(t, "test1", u.Password)
}

func TestFindUserByUsername(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a user in the database
	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")
	assert.Nil(t, err)

	// Try to fetch the user
	u, err := user.FindUserByUsername(tx, "test")
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), u.ID)
	assert.Equal(t, "test", u.Username)
	assert.Equal(t, "test1", u.Password)
}

func TestUpdateUser(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a user in the database
	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")
	assert.Nil(t, err)

	// Fetch the user that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"user\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to update the user
	u := &user.User{
		ID:       id,
		Username: "test2",
		Password: "test",
	}
	err = user.UpdateUser(tx, u)
	assert.Nil(t, err)

	// Check if the user username is changed
	var username string
	var password string
	tx.QueryRow("SELECT * FROM \"user\";").Scan(&id, &username, &password)
	assert.Equal(t, u.Username, username)
	assert.Equal(t, u.Password, password)
}

func TestDeleteUserByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Nil(t, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a user in the database
	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test', 'test1');")
	assert.Nil(t, err)

	// Fetch the user that was just created
	var id uint
	tx.QueryRow("SELECT \"id\" FROM \"user\";").Scan(&id)
	assert.NotEqual(t, uint(0), id)

	// Try to delete the user
	err = user.DeleteUserByID(tx, id)
	assert.Nil(t, err)

	// Check if the user still exists (it should not)
	id = 0
	tx.QueryRow("SELECT \"id\" FROM \"user\";").Scan(&id)
	assert.Equal(t, uint(0), id)
}
