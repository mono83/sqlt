package hooks

import (
	"database/sql"
	"errors"
	"github.com/mono83/sqlt"
	"time"
)

// Querier is a hook container over sqlt.Querier
// capable to run before and after functions
// to modify or log data.
type Querier struct {
	sqlt.Querier

	Before func(string, ...any) (string, []any)
	After  func(*sql.Rows, error, time.Duration, string, ...any) (*sql.Rows, error)
}

// Query is sqlt.Querier interface implementation
func (q Querier) Query(query string, args ...any) (*sql.Rows, error) {
	if q.Querier == nil {
		return nil, errors.New("no hook target for Querier")
	}
	if q.Before != nil {
		query, args = q.Before(query, args...)
	}
	stamp := time.Now()
	rows, err := q.Querier.Query(query, args...)
	if q.After != nil {
		rows, err = q.After(rows, err, time.Since(stamp), query, args...)
	}
	return rows, err
}
