package sqlt

import (
	"database/sql"
	"errors"
)

// ErrCallbackNotSet is returned by CallbackDB when callback for corresponding
// method not set.
var ErrCallbackNotSet = errors.New("callback not set")

// CallbackDB is the simplest implementation of ReaderExecutor interface
// providing to configure which function will be invoked on corresponding
// method.
// This can be useful in unit tests
type CallbackDB struct {
	OnGet    func(dest interface{}, query string, args ...interface{}) error
	OnSelect func(dest interface{}, query string, args ...interface{}) error
	OnExec   func(query string, args ...interface{}) (sql.Result, error)
}

// Get is Getter interface implementation
func (c CallbackDB) Get(dest interface{}, query string, args ...interface{}) error {
	if c.OnGet == nil {
		return ErrCallbackNotSet
	}
	return c.OnGet(dest, query, args...)
}

// Select is Selector interface implementation
func (c CallbackDB) Select(dest interface{}, query string, args ...interface{}) error {
	if c.OnSelect == nil {
		return ErrCallbackNotSet
	}
	return c.OnSelect(dest, query, args...)
}

// Exec is Executor interface implementation
func (c CallbackDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if c.OnExec == nil {
		return nil, ErrCallbackNotSet
	}
	return c.OnExec(query, args...)
}
