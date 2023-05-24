package mysqlt

import (
	"database/sql"
	"testing"

	"github.com/mono83/sqlt"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db := sqlt.CallbackDB{OnExec: func(query string, args ...interface{}) (sql.Result, error) {
		assert.Equal(t, "INSERT INTO `xxx` (`a`,`b`,`c`) VALUES (?,?,?)", query)
		if assert.Len(t, args, 3) {
			assert.Equal(t, 3, args[0])
			assert.Equal(t, "foo", args[1])
			assert.Equal(t, false, args[2])
		}

		return nil, nil
	}}

	_, err := Insert(db, "xxx", map[string]any{
		"a": 3,
		"b": "foo",
		"c": false,
	})
	assert.NoError(t, err)
}
