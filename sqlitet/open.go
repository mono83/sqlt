package sqlitet

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// OpenMem opens connection to in-memory database
// To be used in unit tests
func OpenMem() (*sql.DB, error) {
	return sql.Open("sqlite3", ":memory:")
}

// OpenMemWithMySQLImport opens connection to in-memory database
// and loads SQL files into it. All files are converted from
// MySQL format to SQLite.
//
// Experimental: this feature is experimental and not recommended
// to use due not final API.
func OpenMemWithMySQLImport(files ...string) (*sql.DB, error) {
	db, err := OpenMem()
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		bts, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		s := FromMySQL(string(bts))
		_, err = db.Exec(s)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
