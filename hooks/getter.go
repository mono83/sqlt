package hooks

import (
	"errors"
	"github.com/mono83/sqlt"
	"time"
)

// Getter is a hook container over sqlt.Getter
// capable to run before and after functions
// to modify or log data.
type Getter struct {
	sqlt.Getter

	Before func(any, string, ...any) (any, string, []any)
	After  func(error, time.Duration, any, string, ...any) error
}

// Get is sqlt.Getter interface implementation
func (g Getter) Get(dest any, query string, args ...any) error {
	if g.Getter == nil {
		return errors.New("no hook target for Getter")
	}
	if g.Before != nil {
		dest, query, args = g.Before(dest, query, args...)
	}
	stamp := time.Now()
	err := g.Getter.Get(dest, query, args...)
	if g.After != nil {
		err = g.After(err, time.Since(stamp), dest, query, args...)
	}
	return err
}
