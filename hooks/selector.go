package hooks

import (
	"errors"
	"github.com/mono83/sqlt"
	"time"
)

// Selector is a hook container over sqlt.Selector
// capable to run before and after functions
// to modify or log data.
type Selector struct {
	sqlt.Selector

	Before func(any, string, ...any) (any, string, []any)
	After  func(error, time.Duration, any, string, ...any) error
}

// Select is sqlt.Selector interface implementation
func (s Selector) Select(dest any, query string, args ...any) error {
	if s.Selector == nil {
		return errors.New("no hook target for Selector")
	}
	if s.Before != nil {
		dest, query, args = s.Before(dest, query, args...)
	}
	stamp := time.Now()
	err := s.Selector.Select(dest, query, args...)
	if s.After != nil {
		err = s.After(err, time.Since(stamp), dest, query, args...)
	}
	return err
}
