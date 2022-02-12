package dbTransaction

import (
	"gorm.io/gorm"
)

type rw struct {
	store *gorm.DB
}

func New(db *gorm.DB) rw {
	rw := rw{
		store: db,
	}

	return rw
}

// func (rw rw) WithTx(runner func(tx *sql.Tx) error) error {
// 	tx, err := rw.store.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	err = runner(tx)
// 	if err != nil {
// 		_ = tx.Rollback()
// 		return err
// 	}
// 	return tx.Commit()
// }
