package atomic

import (
	"database/sql"
	"errors"
	"github.com/mono83/sqlt"
)

type transactional interface {
	Begin() (*sql.Tx, error)
}

// Simple performs runs given callback inside SQL transaction.
//
// If given database does not support transactions and mandatory
// flag is set to "true", func will fail with corresponding error.
//
// If given database does not support transactions and mandatory
// flag is set to "false" then callback func will be invoked
// without transaction.
//
// If given callback function fails with error transaction will be
// rolled back, otherwise it will be committed.
func Simple(db sqlt.QuerierExecutor, mandatory bool, f func(sqlt.QuerierExecutor) error) error {
	if db == nil {
		return errors.New("nil db")
	}
	if f == nil {
		return errors.New("nil callback func")
	}

	trx, ok := db.(transactional)
	if !ok {
		// Database connection does not support transactions
		if mandatory {
			return errors.New("unable to start transaction")
		}
		return f(db)
	}

	tx, err := trx.Begin()
	if err != nil {
		return err
	}
	if err := f(tx); err != nil {
		// Error invoking callback function, rolling back
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
