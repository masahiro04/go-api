package dbTransaction

import (
	"database/sql"

	uc "clean_architecture/golang/usecases"
)

type rw struct {
	store *sql.DB
}

func New(db *sql.DB) uc.DBTransaction {
	rw := rw{
		store: db,
	}

	return rw
}

func (rw rw) WithTx(runner func(tx *sql.Tx) error) error {
	tx, err := rw.store.Begin()
	if err != nil {
		return err
	}
	err = runner(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
