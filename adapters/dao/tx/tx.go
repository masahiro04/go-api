package tx

import (
	"gorm.io/gorm"
)

type rw struct {
	db *gorm.DB
}

func New(db *gorm.DB) *rw {
	return &rw{
		db: db,
	}
}

func (rw rw) WithTx(runner func(tx *gorm.DB) error) error {
	tx := rw.db.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	if err := runner(rw.db); err != nil {
		tx.Rollback()
		// tx.Errorだとnilが変える可能性あるのでerrで対応
		return err
	}
	return tx.Commit().Error
}
