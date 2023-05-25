package mysqlt

import (
	"database/sql"
	"testing"

	"github.com/mono83/sqlt"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db := sqlt.CallbackDB{OnExec: func(query string, args ...interface{}) (sql.Result, error) {
		if query == "INSERT INTO `xxx` (`a`,`b`) VALUES (?,?)" {
			if assert.Len(t, args, 2) {
				assert.Equal(t, 3, args[0])
				assert.Equal(t, "foo", args[1])
			}
		} else if query == "INSERT INTO `xxx` (`b`,`a`) VALUES (?,?)" {
			if assert.Len(t, args, 2) {
				assert.Equal(t, 3, args[1])
				assert.Equal(t, "foo", args[0])
			}
		} else {
			t.Fail()
		}

		return nil, nil
	}}

	_, err := Insert(db, "xxx", map[string]any{
		"a": 3,
		"b": "foo",
	})
	assert.NoError(t, err)
}
