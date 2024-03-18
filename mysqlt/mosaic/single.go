package mosaic

import (
	"database/sql"
	"github.com/mono83/sqlt"
)

func MakePersistentSingleWriter[T any](
	db sqlt.QuerierExecutor,
	table string,
	transactional bool,
	columns []string,
	mapping func(T) []any,
) SingleWriter[T] {
	list := MakePersistentListWriter(db, table, transactional, columns, mapping)
	return func(primaryID, typeID uint64, data T) error {
		return list(primaryID, typeID, []T{data})
	}
}

func MakePersistentSingeReader[T any](
	db sqlt.Querier,
	table string,
	columns []string,
	mapping func([]any) (*T, error),
) SingeReader[T] {
	list := MakePersistentListReader(db, table, columns, mapping)
	return func(primaryID, typeID uint64) (*T, error) {
		values, err := list(primaryID, typeID)
		if err != nil {
			return nil, err
		}
		if len(values) == 0 {
			return nil, sql.ErrNoRows
		}
		return &values[0], nil
	}
}
