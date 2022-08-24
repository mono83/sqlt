package sqlt

import "database/sql"

// Executor is a thin interface for database connection capable to modify data
type Executor interface {
	Exec(query string, args ...any) (sql.Result, error)
}

// Selector is a thin interface for database connection capable to read data
type Selector interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// SelectorExecutor is a thin interface for database connection capable to both read and write data
type SelectorExecutor interface {
	Selector
	Executor
}
