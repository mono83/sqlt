package mysqlt

import (
	"testing"

	"github.com/mono83/sqlt"
	"github.com/stretchr/testify/assert"
)

func TestSelectByID(t *testing.T) {
	db := sqlt.CallbackDB{OnSelect: func(_ interface{}, query string, args ...interface{}) error {
		assert.Equal(t, "SELECT * FROM users WHERE `id` IN (?,?)", query)
		if assert.Len(t, args, 2) {
			assert.Equal(t, 21, args[0])
			assert.Equal(t, 88, args[1])
		}

		return nil
	}}

	_, err := SelectByID[user](db, "users", 21, 88)
	assert.NoError(t, err)
}

func TestSelectByID1(t *testing.T) {
	db := sqlt.CallbackDB{OnSelect: func(_ interface{}, query string, args ...interface{}) error {
		assert.Equal(t, "SELECT * FROM users WHERE `id`=?", query)
		if assert.Len(t, args, 1) {
			assert.Equal(t, 6, args[0])
		}

		return nil
	}}

	_, err := SelectByID[user](db, "users", 6)
	assert.NoError(t, err)
}

func TestMakeSelectByID(t *testing.T) {
	db := sqlt.CallbackDB{OnSelect: func(_ interface{}, query string, args ...interface{}) error {
		assert.Equal(t, "SELECT * FROM blocked WHERE `id` IN (?,?,?)", query)
		if assert.Len(t, args, 3) {
			assert.Equal(t, "foo", args[0])
			assert.Equal(t, "bar", args[1])
			assert.Equal(t, "baz", args[2])
		}

		return nil
	}}

	sel := MakeSelectByID[user, string](db, "blocked")
	_, err := sel("foo", "bar", "baz")
	assert.NoError(t, err)
}
