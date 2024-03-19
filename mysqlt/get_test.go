package mysqlt

import (
	"testing"

	"github.com/mono83/sqlt"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	db := sqlt.CallbackDB{OnGet: func(_ any, query string, args ...any) error {
		assert.Equal(t, "SELECT * FROM `users` WHERE `id`=?", query)
		if assert.Len(t, args, 1) {
			assert.Equal(t, 12, args[0])
		}

		return nil
	}}

	_, err := GetByID[user](db, "users", 12)
	assert.NoError(t, err)
}

func TestMakeGetByID(t *testing.T) {
	db := sqlt.CallbackDB{OnGet: func(_ any, query string, args ...any) error {
		assert.Equal(t, "SELECT * FROM `blocked` WHERE `id`=?", query)
		if assert.Len(t, args, 1) {
			assert.Equal(t, "who", args[0])
		}

		return nil
	}}

	getter := MakeGetByID[user, string](db, "blocked")
	_, err := getter("who")
	assert.NoError(t, err)
}

type user struct {
}
