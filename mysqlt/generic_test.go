package mysqlt

import (
	"github.com/jmoiron/sqlx"
	"github.com/mono83/sqlt"
	"github.com/mono83/sqlt/sqlitet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeGenericReader(t *testing.T) {
	db, err := sqlitet.OpenMemWithMySQLImport(
		"../.provision/tests/context.sql",
		"../.provision/tests/context-data.sql",
	)
	defer db.Close()
	if assert.NoError(t, err) {
		xdb := sqlx.NewDb(db, "sqlite3")
		read := MakeGenericReader[uint64, genericContext](xdb, "context", "id")
		if c, err := read(2); assert.NoError(t, err) {
			assert.Equal(t, uint64(2), c.ID)
			assert.Equal(t, "foo", c.UUID)
			assert.Equal(t, uint32(2356372769), c.UUIDHash)
			assert.Equal(t, int64(123456789), c.CreatedAt.Time().Unix())
		}
	}
}

type genericContext struct {
	ID        uint64           `db:"id"`
	UUID      string           `db:"uuid"`
	UUIDHash  uint32           `db:"uuidHash"`
	CreatedAt sqlt.UnixSeconds `db:"createdAt"`
}
