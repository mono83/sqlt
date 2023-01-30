package sqlt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeStdGetterByColumn(t *testing.T) {
	db := CallbackDB{OnGet: func(_ interface{}, query string, args ...interface{}) error {
		assert.Equal(t, "SELECT * FROM users WHERE id = ?", query)
		if assert.Len(t, args, 1) {
			assert.Equal(t, int64(12), args[0])
		}

		return nil
	}}

	get := MakeStdGetterByColumn[int64, user](db, "users", "id")
	if _, err := get(12); assert.NoError(t, err) {
	}
}

func TestMakeStdSelectorByColumn(t *testing.T) {
	db := CallbackDB{OnSelect: func(_ interface{}, query string, args ...interface{}) error {
		assert.Equal(t, "SELECT * FROM settings WHERE userId IN (?,?)", query)
		if assert.Len(t, args, 2) {
			assert.Equal(t, int64(9), args[0])
			assert.Equal(t, int64(-2), args[1])
		}
		return nil
	}}

	sel := MakeStdSelectorByColumn[int64, user](db, "settings", "userId")
	if _, err := sel(9, -2); assert.NoError(t, err) {
	}
}

type user struct {
}
