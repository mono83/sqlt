package sqlt

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeStdGetterByColumn(t *testing.T) {
	db := &singleMockDb{}

	get := MakeStdGetterByColumn[int64, user](db, "users", "id")
	if _, err := get(12); assert.Equal(t, mockErr, err) {
		assert.Equal(t, "SELECT * FROM users WHERE id = ?", db.query)
		if assert.Len(t, db.args, 1) {
			assert.Equal(t, int64(12), db.args[0])
		}
	}
}

func TestMakeStdSelectorByColumn(t *testing.T) {
	db := &singleMockDb{}

	sel := MakeStdSelectorByColumn[int64, user](db, "settings", "userId")
	if _, err := sel(9, -2); assert.Equal(t, mockErr, err) {
		assert.Equal(t, "SELECT * FROM settings WHERE userId IN (?,?)", db.query)
		if assert.Len(t, db.args, 2) {
			assert.Equal(t, int64(9), db.args[0])
			assert.Equal(t, int64(-2), db.args[1])
		}
	}
}

var mockErr = errors.New("mock")

type singleMockDb struct {
	query string
	args  []any
}

type user struct {
}

func (s *singleMockDb) Get(_ interface{}, query string, args ...interface{}) error {
	s.query = query
	s.args = args
	return mockErr
}
