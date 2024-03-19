package sqlt

import "database/sql"

// Rowset is plain two-dimensional data table
// containing data, obtained from SQL database
// in non-structured way (via interface{})
type Rowset struct {
	ColumnNames []string
	ColumnTypes []*sql.ColumnType
	Rows        [][]any
}

// Clear cleans up rowset contents.
func (r *Rowset) Clear() {
	r.ColumnNames = nil
	r.ColumnTypes = nil
	r.Rows = nil
}

// Scan method fills rowset contents using data
// from given sql.Rows.
func (r *Rowset) Scan(rows *sql.Rows) error {
	// Clearing current state
	r.Clear()

	return IterateScan(rows, func(names []string, types []*sql.ColumnType, row []any) {
		if r.ColumnNames == nil {
			// First row
			r.ColumnNames = names
			r.ColumnTypes = types
		}
		r.Rows = append(r.Rows, row)
	})
}

// Size returns amount of rows.
func (r Rowset) Size() int {
	return len(r.Rows)
}

// Each iterated over each row, passing it with
// corresponding metadata to callback function.
func (r Rowset) Each(f func([]string, []*sql.ColumnType, []any)) {
	if f != nil {
		for _, row := range r.Rows {
			f(r.ColumnNames, r.ColumnTypes, row)
		}
	}
}

// SliceMap returns content of rowset as slice of maps.
func (r Rowset) SliceMap() (out []map[string]any) {
	cnt := len(r.ColumnNames)
	r.Each(func(names []string, _ []*sql.ColumnType, data []any) {
		m := make(map[string]any, cnt)
		for i := 0; i < cnt; i++ {
			m[names[i]] = data[i]
		}
		out = append(out, m)
	})
	return
}

// MapValuesStd maps all values in rowset using standard conversion function.
func (r Rowset) MapValuesStd() Rowset {
	return r.MapValues(StdConvert)
}

// MapValues maps all values in rowset using given conversion function.
func (r Rowset) MapValues(f func(src any, def *sql.ColumnType) any) Rowset {
	if f == nil {
		return r
	}
	out := Rowset{
		ColumnNames: r.ColumnNames,
		ColumnTypes: r.ColumnTypes,
	}

	cnt := len(r.ColumnNames)
	for _, src := range r.Rows {
		row := make([]any, cnt)
		for i := 0; i < cnt; i++ {
			row[i] = f(src[i], r.ColumnTypes[i])
		}
		out.Rows = append(out.Rows, row)
	}
	return out
}
