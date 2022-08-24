package sqlt

import "database/sql"

// InsertID returns auto incremental identifier of last
// inserted entity as int64.
func InsertID(result sql.Result, err error) (int64, error) {
	if err != nil {
		// Received error in argument
		return 0, err
	}

	return result.LastInsertId()
}

// InsertIDU returns auto incremental identifier of last
// inserted entity as uint64.
func InsertIDU(result sql.Result, err error) (uint64, error) {
	id, err := InsertID(result, err)
	if err == nil {
		return uint64(id), nil
	}
	return 0, err
}
