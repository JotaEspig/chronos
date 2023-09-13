package tests

import (
	"chronos/config"
	"chronos/pkg/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUserByID(t *testing.T) {
	tx, err := config.DB.Begin()
	assert.Equal(t, nil, err)
	defer tx.Rollback()

	cleanDB(tx)
	tx.Exec("INSERT INTO \"user\" VALUES (1, 'test')")
	u := user.FindUserByID(1)
	assert.NotEqual(t, nil, u)
	return
}
