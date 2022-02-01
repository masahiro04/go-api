package userDao_test

import (
	"testing"
)

func TestRw_happyCreate(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()
	//
	// rw := userDao.New(db)
	// blog := testData.Blog()
	// //id := 1
	// mock.ExpectQuery(regexp.QuoteMeta(userDao.CreateSql)).
	// 	WithArgs(blog.Title.Value, blog.Body.Value, AnyTime{}, AnyTime{}).
	// 	WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(blog.ID.Value))
	//
	// if _, err = rw.Create(blog); err != nil {
	// 	t.Errorf("error was not expected while updating stats: %s", err)
	// }
	//
	// // 使用されたモックDBが期待通りの値を持っているかを検証
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
}

func TestRw_unHappyCreate(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()
	//
	// rw := userDao.New(db)
	// blog := testData.Blog()
	//
	// mock.ExpectQuery(regexp.QuoteMeta(userDao.CreateSql)).
	// 	WillReturnError(fmt.Errorf("some error"))
	//
	// if _, err = rw.Create(blog); err == nil {
	// 	t.Errorf("error was not expected while updating stats: %s", err)
	// }
	//
	// // err == nil -> つまり、エラーがあればpassするということ
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
}
