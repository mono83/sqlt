package sqlt

import "bytes"

// MakeStdGetterByColumn creates function that will fetch single entity
// from database by `column=value` criteria.
// Will return sql.ErrNoRows if no matches and driver-specific error if
// more than one entity matched.
// To match more than one, use MakeStdSelectorByColumn.
func MakeStdGetterByColumn[K any, V any](db Getter, table, column string) func(key K) (*V, error) {
	query := "SELECT * FROM " + table + " WHERE " + column + " = ?"
	return func(key K) (*V, error) {
		var target V
		if err := db.Get(&target, query, key); err != nil {
			return nil, err
		}
		return &target, nil
	}
}

// MakeStdSelectorByColumn created function that will fetch multiple entities
// from database by `column IN (keys...)` criteria.
func MakeStdSelectorByColumn[K any, V any](db Selector, table, column string) func(keys ...K) ([]V, error) {
	queryPrefix := "SELECT * FROM " + table + " WHERE " + column + " IN ("
	return func(keys ...K) ([]V, error) {
		if len(keys) == 0 {
			return nil, nil // No data and no error
		}
		ikeys := make([]any, len(keys))
		query := bytes.NewBufferString(queryPrefix)
		for i, j := range keys {
			ikeys[i] = j
			if i == 0 {
				query.WriteRune('?')
			} else {
				query.WriteString(",?")
			}
		}
		query.WriteRune(')')

		var target []V
		if err := db.Select(&target, query.String(), ikeys...); err != nil {
			return nil, err
		}
		return target, nil
	}
}
