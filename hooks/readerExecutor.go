package hooks

import "database/sql"

// ReaderExecutor is a hook container over sqlt.ReaderExecutor
// capable to run before and after functions
// to modify or log data.
type ReaderExecutor struct {
	Getter
	Selector
	Executor
}

// Get is sqlt.Getter and sqlt.ReaderExecutor interfaces implementation
func (r ReaderExecutor) Get(dest any, query string, args ...any) error {
	return r.Getter.Get(dest, query, args...)
}

// Select is sqlt.Selector and sqlt.ReaderExecutor interfaces implementation
func (r ReaderExecutor) Select(dest any, query string, args ...any) error {
	return r.Selector.Select(dest, query, args...)
}

// Exec is sqlt.Executor and sqlt.ReaderExecutor interfaces implementation
func (r ReaderExecutor) Exec(query string, args ...any) (sql.Result, error) {
	return r.Executor.Exec(query, args...)
}
