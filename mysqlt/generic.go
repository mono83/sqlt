package mysqlt

import (
	"database/sql"
	"errors"
	"github.com/mono83/sqlt"
)

// MakeGenericListReader constructs accessor to database table
// able to match records by single column value (like id).
func MakeGenericListReader[I, T any](db sqlt.Reader, table, column string) sqlt.GenericListReader[I, T] {
	query := "SELECT * FROM " + Name(table) + " WHERE " + Name(column) + " IN ("
	return func(ids ...I) (out []T, err error) {
		l := len(ids)
		if l == 0 {
			return nil, sql.ErrNoRows
		}

		args := make([]any, l, l)
		for i, j := range ids {
			args[i] = j
		}

		err = db.Select(&out, query+sqlt.PlaceholdersString(l)+")", args...)
		return
	}
}

// MakeGenericReader constructs accessor to database table
// able to match single record by single column value (like id).
func MakeGenericReader[I, T any](db sqlt.Reader, table, column string) sqlt.GenericReader[I, T] {
	list := MakeGenericListReader[I, T](db, table, column)
	return func(ID I) (*T, error) {
		data, err := list(ID)
		if err != nil {
			return nil, err
		}
		if l := len(data); l == 0 {
			return nil, sql.ErrNoRows
		} else if l > 1 {
			return nil, errors.New("to many rows")
		}
		return &data[0], nil
	}
}
