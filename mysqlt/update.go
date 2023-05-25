package mysqlt

import (
	"bytes"
	"database/sql"
	"errors"
	"github.com/mono83/sqlt"
)

// Update performs database modification query
func Update(exec sqlt.Executor, table string, values map[string]any, condition string, iargs ...any) (sql.Result, error) {
	if len(values) == 0 {
		return nil, errors.New("no values to update")
	}

	var args []any
	query := bytes.NewBufferString("UPDATE ")
	query.WriteString(Name(table))
	query.WriteString(" SET ")
	for k, v := range values {
		args = append(args, v)
		if len(args) > 1 {
			query.WriteRune(',')
		}
		query.WriteString(Name(k))
		query.WriteString("=?")
	}
	query.WriteString(" WHERE ")
	query.WriteString(condition)
	args = append(args, iargs...)

	return exec.Exec(query.String(), args...)
}
