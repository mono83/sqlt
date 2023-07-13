package mysqlt

import (
	"bytes"
	"database/sql"
	"errors"

	"github.com/mono83/sqlt"
)

// Insert places values map into database
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

// MakeInsert constructs function that places values into database
func MakeInsert[IN any, OUT any](
	exec sqlt.Executor,
	table string,
	mapIn func(IN) (map[string]any, error),
	mapOut func(sql.Result) (OUT, error),
) func(IN) (OUT, error) {
	return func(in IN) (o OUT, err error) {
		// Converting value to save
		values, err := mapIn(in)
		if err == nil {
			var res sql.Result
			res, err = Insert(exec, table, values)
			if err == nil {
				o, err = mapOut(res)
			}
		}
		return
	}
}
