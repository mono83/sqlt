package sqlt

import "database/sql"

// IterateScan scans given rows ony by one passing obtained data to
// callback function.
func IterateScan(rows *sql.Rows, callback func([]string, []*sql.ColumnType, []any)) error {
	if rows == nil {
		return sql.ErrNoRows
	}
	if callback == nil {
		return nil // No error
	}

	// Reading column names
	columnNames, err := rows.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	columnCount := len(columnNames)

	processed := false
	for rows.Next() {
		// Building columns and their pointers for
		// sql.Scan func.
		columns := make([]any, columnCount)
		columnPointers := make([]any, columnCount)
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scanning the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return err
		}

		callback(columnNames, columnTypes, columns)
		processed = true
	}
	if !processed {
		return sql.ErrNoRows
	}
	return nil
}

// IterateScanE scans given rows ony by one passing obtained data to
// callback function. On any error stops iteration and returns it.
func IterateScanE(rows *sql.Rows, callback func([]string, []*sql.ColumnType, []any) error) error {
	if rows == nil {
		return sql.ErrNoRows
	}
	if callback == nil {
		return nil // No error
	}

	// Reading column names
	columnNames, err := rows.Columns()
	if err != nil {
		return err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	columnCount := len(columnNames)

	processed := false
	for rows.Next() {
		// Building columns and their pointers for
		// sql.Scan func.
		columns := make([]any, columnCount)
		columnPointers := make([]any, columnCount)
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scanning the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return err
		}

		err = callback(columnNames, columnTypes, columns)
		if err != nil {
			return err
		}
		processed = true
	}
	if !processed {
		return sql.ErrNoRows
	}
	return nil
}
