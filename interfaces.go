package sqlt

import "database/sql"

// Getter is a thin interface for database connection capable to read data
type Getter interface {
	Get(dest interface{}, query string, args ...interface{}) error
}

// Selector is a thin interface for database connection capable to read data
type Selector interface {
	Select(dest interface{}, query string, args ...interface{}) error
}

// Reader is an interface defining connection able to perform read operations
type Reader interface {
	Getter
	Selector
}

// Executor is a thin interface for database connection capable to modify data
type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// ReaderExecutor is a thin interface for database connection capable to both read and write data
type ReaderExecutor interface {
	Reader
	Executor
}
