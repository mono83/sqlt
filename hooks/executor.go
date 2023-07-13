package hooks

import (
	"database/sql"
	"errors"
	"github.com/mono83/sqlt"
	"time"
)

// Executor is a hook container over sqlt.Executor
// capable to run before and after functions
// to modify or log data.
type Executor struct {
	sqlt.Executor

	Before func(string, ...any) (string, []any)
	After  func(sql.Result, error, time.Duration, string, ...any) (sql.Result, error)
}

// Exec is sqlt.Executor interface implementation
func (e Executor) Exec(query string, args ...any) (sql.Result, error) {
	if e.Executor == nil {
		return nil, errors.New("no hook target for Executor")
	}
	if e.Before != nil {
		query, args = e.Before(query, args...)
	}
	stamp := time.Now()
	res, err := e.Executor.Exec(query, args...)
	if e.After != nil {
		res, err = e.After(res, err, time.Since(stamp), query, args...)
	}
	return res, err
}
