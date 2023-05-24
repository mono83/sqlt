package mysqlt

import (
	"github.com/mono83/sqlt"
)

// SelectByID reads multiple entities by identifiers
func SelectByID[T any](selector sqlt.Selector, table string, id ...any) (out []T, err error) {
	if len(id) > 0 {
		sql, args := In(id...)
		err = selector.Select(&out, "SELECT * FROM "+table+" WHERE `id`"+sql, args...)
	}
	return
}

// MakeSelectByID constructs function that reads multiple entities by identifiers
func MakeSelectByID[T any, I any](selector sqlt.Selector, table string) func(...I) ([]T, error) {
	return func(i ...I) ([]T, error) {
		return SelectByID[T](selector, table, drain(i)...)
	}
}
