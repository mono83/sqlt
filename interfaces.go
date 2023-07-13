package sqlt

import "database/sql"

// Querier is a thin interface for database connection capable to query data
type Querier interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

// Getter is a thin interface for database connection capable to read data
type Getter interface {
	Get(dest any, query string, args ...any) error
}

// Selector is a thin interface for database connection capable to read data
type Selector interface {
	Select(dest any, query string, args ...any) error
}

// Reader is an interface defining connection able to perform read operations
type Reader interface {
	Getter
	Selector
}

// Executor is a thin interface for database connection capable to modify data
type Executor interface {
	Exec(query string, args ...any) (sql.Result, error)
}

// ReaderExecutor is a thin interface for database connection capable to both read and write data
type ReaderExecutor interface {
	Reader
	Executor
}
