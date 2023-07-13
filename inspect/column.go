package inspect

// Column is high level abstraction over database column
type Column struct {
	Database string
	Table    string
	Name     string
	Comment  string

	Type     Type
	Nullable bool
	Unsigned bool
	Indexed  bool
	Size     int
	Values   []string

	OriginalType string
}
