package tests

import (
	"chronos/config"
	"chronos/pkg/models/user"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByID(t *testing.T) {
	godotenv.Load("../.env")
	config.InitDB()

	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)
	_, err = tx.Exec("INSERT INTO \"user\" VALUES (1, 'test');")
	assert.Equal(t, nil, err)

	u := user.FindUserByID(tx, 1)
	assert.NotEqual(t, nil, u)
	assert.Equal(t, uint(1), u.ID)
	assert.Equal(t, "test", u.Username)

	return
}
