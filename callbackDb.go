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
	OnGet    func(dest any, query string, args ...any) error
	OnSelect func(dest any, query string, args ...any) error
	OnExec   func(query string, args ...any) (sql.Result, error)
}

// Get is Getter interface implementation
func (c CallbackDB) Get(dest any, query string, args ...any) error {
	if c.OnGet == nil {
		return ErrCallbackNotSet
	}
	return c.OnGet(dest, query, args...)
}

// Select is Selector interface implementation
func (c CallbackDB) Select(dest any, query string, args ...any) error {
	if c.OnSelect == nil {
		return ErrCallbackNotSet
	}
	return c.OnSelect(dest, query, args...)
}

// Exec is Executor interface implementation
func (c CallbackDB) Exec(query string, args ...any) (sql.Result, error) {
	if c.OnExec == nil {
		return nil, ErrCallbackNotSet
	}
	return c.OnExec(query, args...)
}
