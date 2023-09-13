package tests

import (
	"chronos/config"
	"chronos/pkg/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	newUser := &user.User{
		Username: "newuser",
	}

	err = user.CreateUser(tx, newUser)
	assert.Equal(t, nil, err)

	// Check that the user has been successfully created and has a non-zero ID.
	assert.NotEqual(t, uint(0), newUser.ID)

	// Fetch the user from the database by ID and check if it matches the created user.
	var id uint
	var username string

	// Check if it was created
	tx.QueryRow("SELECT \"id\", \"username\" FROM \"user\";").Scan(&id, &username)
	assert.Equal(t, nil, err)
	assert.Equal(t, newUser.ID, id)
	assert.Equal(t, newUser.Username, username)
}

func TestFindUserByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	// Insert a user in the database
	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")
	assert.Equal(t, nil, err)

	// Try to fetch the user
	u, err := user.FindUserByID(tx, 1)
	assert.NotEqual(t, nil, u)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint(1), u.ID)
	assert.Equal(t, "test", u.Username)
}

func TestFindUserByUsername(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)

	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")
	assert.Equal(t, nil, err)

	u, err := user.FindUserByUsername(tx, "test")
	assert.NotEqual(t, nil, u)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint(1), u.ID)
	assert.Equal(t, "test", u.Username)
}
