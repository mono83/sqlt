package hooks

// Reader is a hook container over sqlt.Reader
// capable to run before and after functions
// to modify or log data.
type Reader struct {
	Getter
	Selector
}

// Get is sqlt.Getter and sqlt.Reader interfaces implementation
func (r Reader) Get(dest any, query string, args ...any) error {
	return r.Getter.Get(dest, query, args...)
}

// Select is sqlt.Selector and sqlt.Reader interfaces implementation
func (r Reader) Select(dest any, query string, args ...any) error {
	return r.Selector.Select(dest, query, args...)
}
