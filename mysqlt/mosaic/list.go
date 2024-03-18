package mosaic

import (
	"bytes"
	"database/sql"
	"errors"
	"github.com/mono83/sqlt"
	"github.com/mono83/sqlt/atomic"
	"github.com/mono83/sqlt/mysqlt"
	"strings"
	"time"
)

func MakePersistentListReader[T any](
	db sqlt.Querier,
	table string,
	columns []string,
	mapping func([]any) (*T, error),
) ListReader[T] {
	// Building select statement
	b := bytes.NewBufferString("SELECT ")
	for i, col := range columns {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(mysqlt.Name(col))
	}
	b.WriteString(" FROM ")
	b.WriteString(mysqlt.Name(table))
	b.WriteString(" WHERE `primaryId`=? AND `typeId`=? AND `enabled`=? LIMIT 1")
	sqlSelect := b.String()
	return func(primaryID, typeID uint64) ([]T, error) {
		rows, err := db.Query(sqlSelect, primaryID, typeID, "true")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var out []T
		err = sqlt.IterateScanE(rows, func(columns []string, types []*sql.ColumnType, data []interface{}) error {
			x, err := mapping(data)
			if err == nil {
				out = append(out, *x)
			}
			return err
		})
		if err != nil {
			return nil, err
		}
		if out == nil {
			return nil, sql.ErrNoRows
		}
		return out, nil
	}
}

func MakePersistentListWriter[T any](
	db sqlt.QuerierExecutor,
	table string,
	transactional bool,
	columns []string,
	mapping func(T) []any,
) ListWriter[T] {
	// Building insert statement
	insertBuf := bytes.NewBufferString("INSERT INTO ")
	repeatBuf := bytes.NewBufferString(",(")

	insertBuf.WriteString(mysqlt.Name(table))
	insertBuf.WriteString(" (`primaryId`,`typeId`,`createdAt`,`enabled`")
	for _, col := range columns {
		insertBuf.WriteString(",")
		insertBuf.WriteString(mysqlt.Name(col))
	}
	insertBuf.WriteString(") VALUES (")
	insertBuf.WriteString(sqlt.PlaceholdersString(4 + len(columns)))
	repeatBuf.WriteString(sqlt.PlaceholdersString(4 + len(columns)))
	insertBuf.WriteString(")")
	repeatBuf.WriteString(")")
	sqlInsert := insertBuf.String()
	sqlRepeat := repeatBuf.String()

	// Building update statement
	sqlDelete := "UPDATE " + mysqlt.Name(table) + " SET `enabled`=? WHERE `id` IN "

	// Building select statement
	selectBuf := bytes.NewBufferString("SELECT `id`")
	for _, col := range columns {
		selectBuf.WriteString(",")
		selectBuf.WriteString(mysqlt.Name(col))
	}
	selectBuf.WriteString(" FROM ")
	selectBuf.WriteString(mysqlt.Name(table))
	selectBuf.WriteString(" WHERE `primaryId`=? AND `typeId`=? AND `enabled`=?")
	if transactional {
		selectBuf.WriteString(" FOR UPDATE")
	}
	sqlSelect := selectBuf.String()

	return func(primaryID, typeID uint64, data []T) error {
		if len(data) == 0 {
			return nil
		}

		now := time.Now().Unix()

		// Mapping all
		var allMapped [][]any
		for _, datum := range data {
			mapped := mapping(datum)
			if len(mapped) != len(columns) {
				return errors.New("mapped columns count does not match expected")
			}
			allMapped = append(allMapped, mapped)
		}

		return atomic.Simple(db, transactional, func(db sqlt.QuerierExecutor) error {
			// Reading current data
			var current []existing
			rows, err := db.Query(sqlSelect, primaryID, typeID, "true")
			if err != nil {
				return err
			}
			defer rows.Close()
			err = sqlt.IterateScanE(rows, func(_ []string, _ []*sql.ColumnType, data []interface{}) error {
				current = append(current, existing{ID: uint64(data[0].(int64)), Data: data[1:]})
				return nil
			})
			if err != nil {
				return err
			}

			// Detecting changes
			_, deleted, pending := detectChanges(current, allMapped)

			// Inserting new
			if len(pending) > 0 {
				query := sqlInsert + strings.Repeat(sqlRepeat, len(pending)-1)
				var args []any
				for _, x := range pending {
					args = append(args, primaryID, typeID, now, "true")
					args = append(args, x...)
				}

				_, err := db.Exec(query, args...)
				if err != nil {
					return err
				}
			}

			// Deleting previous
			if len(deleted) > 0 {
				query := sqlDelete + "(" + sqlt.PlaceholdersString(len(deleted)) + ")"
				var args []any
				args = append(args, "false")
				for _, x := range deleted {
					args = append(args, x)
				}
				_, err := db.Exec(query, args...)
				if err != nil {
					return err
				}
			}

			// TODO update

			return err
		})
	}
}
