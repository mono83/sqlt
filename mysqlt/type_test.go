package mysqlt

import (
	"github.com/jmoiron/sqlx"
	"github.com/mono83/sqlt"
	"github.com/mono83/sqlt/sqlitet"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCustomDataTypes(t *testing.T) {
	db, err := sqlitet.OpenMemWithMySQLImport(
		"../.provision/tests/customDataTypes.sql",
	)
	if assert.NoError(t, err) {
		_, err = db.Exec(
			"INSERT INTO `customDataTypes` VALUES (null, 1711117434, -1711117434, 1711117434187, -1711117434187, 123456789, 'true','{\"type\":8,\"value\":\"foo\"}')",
		)
		if assert.NoError(t, err) {
			dbx := sqlx.NewDb(db, "sqlite3")
			var e typeTestEntity
			err = dbx.Get(&e, "SELECT * FROM `customDataTypes`")
			if assert.NoError(t, err) {
				assert.Equal(t, uint64(1), e.ID)
				assert.Equal(t, time.Date(2024, time.March, 22, 14, 23, 54, 0, time.UTC), e.USU.Time())
				assert.Equal(t, time.Date(1915, time.October, 12, 9, 36, 6, 0, time.UTC), e.USS.Time())
				assert.Equal(t, time.Date(2024, time.March, 22, 14, 23, 54, 187000000, time.UTC), e.UMU.Time())
				assert.Equal(t, time.Date(1915, time.October, 12, 9, 36, 5, 813000000, time.UTC), e.UMS.Time())
				assert.Equal(t, time.Nanosecond*123456789, e.DN.Duration())
				assert.True(t, e.TF.Bool())

				var se typeTestSubentity
				if err := e.Json.Unmarshal(&se); assert.NoError(t, err) {
					assert.Equal(t, 8, se.Type)
					assert.Equal(t, "foo", se.Value)
				}
			}
		}
	}
}

type typeTestEntity struct {
	ID   uint64             `db:"id"`
	USU  sqlt.UnixSeconds   `db:"typeUnixSeconds"`
	USS  sqlt.UnixSeconds   `db:"typeUnixSecondsSigned"`
	UMU  sqlt.UnixMillis    `db:"typeUnixMillis"`
	UMS  sqlt.UnixMillis    `db:"typeUnixMillisSigned"`
	DN   sqlt.DurationNanos `db:"typeElapsedNanos"`
	TF   sqlt.TrueFalse     `db:"enabled"`
	Json sqlt.JSON          `db:"json"`
}

type typeTestSubentity struct {
	Type  int    `json:"type"`
	Value string `json:"value"`
}
