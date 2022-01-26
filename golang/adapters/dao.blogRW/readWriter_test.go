package blogRW_test

// AnyTimeはテストを通すための設定。公式を参考に実装
// type AnyTime struct{}
//
// func (a AnyTime) Match(v driver.Value) bool {
// 	_, ok := v.(time.Time)
// 	return ok
// }

//

//
// func TestRw_happyDelete(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//
// 	rw := blogRW.New(db)
// 	blog := testData.Blog()
//
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE blogs SET updated_at = $2, deleted_at = $3 WHERE id = $1`)).
// 		WithArgs(blog.ID, AnyTime{}, AnyTime{}).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
//
// 	// モック化されたDBを用いてテスト対象関数を実行
// 	if err = rw.Delete(blog.ID.Value()); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
//
// 	// 使用されたモックDBが期待通りの値を持っているかを検証
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
//
// func TestRw_unHappyDelete(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//
// 	rw := blogRW.New(db)
// 	blog := testData.Blog()
//
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE blogs SET updated_at = $2, deleted_at = $3 WHERE id = $1`)).
// 		WithArgs(blog.ID, AnyTime{}, AnyTime{}).
// 		WillReturnError(fmt.Errorf("some error"))
//
// 	// モック化されたDBを用いてテスト対象関数を実行
// 	if err = rw.Delete(blog.ID.Value()); err == nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
//
// 	// 使用されたモックDBが期待通りの値を持っているかを検証
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
