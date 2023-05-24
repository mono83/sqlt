package mysqlt

import (
	"bytes"
	"database/sql"
	"errors"

	"github.com/mono83/sqlt"
)

func Insert(exec sqlt.Executor, table string, values map[string]any) (sql.Result, error) {
	if len(values) == 0 {
		return nil, errors.New("no values to insert")
	}

	var args []any
	query := bytes.NewBufferString("INSERT INTO ")
	query.WriteString(Name(table))
	query.WriteString(" (")
	for k, v := range values {
		args = append(args, v)
		if len(args) > 1 {
			query.WriteRune(',')
		}
		query.WriteString(Name(k))
	}
	query.WriteString(") VALUES (")
	for i := range args {
		if i > 0 {
			query.WriteRune(',')
		}
		query.WriteRune('?')
	}
	query.WriteRune(')')

	return exec.Exec(query.String(), args...)
}
