package tx

// type rw struct {
// 	db *gorm.DB
// }
//
// func New(db *gorm.DB) *rw {
// 	return &rw{
// 		db: db,
// 	}
// }
//
// func (rw rw) WithTx(runner func(tx rw) error) error {
// 	tx := rw.db.Begin()
// 	if tx.Error != nil {
// 		tx.Rollback()
// 		return tx.Error
// 	}
//
// 	if err := runner(rw); err != nil {
// 		tx.Rollback()
// 		return tx.Error
// 	}
//
// 	return tx.Commit().Error
// }
