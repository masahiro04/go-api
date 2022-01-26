package blogRW_test

// AnyTimeはテストを通すための設定。公式を参考に実装
// type AnyTime struct{}
//
// func (a AnyTime) Match(v driver.Value) bool {
// 	_, ok := v.(time.Time)
// 	return ok
// }

//
// func TestRw_happyCreate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//
// 	rw := blogRW.New(db)
// 	blog := testData.Blog()
// 	//id := 1
// 	mock.ExpectQuery(regexp.QuoteMeta(
// 		`INSERT INTO blogs (title, body, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING id`)).
// 		WithArgs(blog.Title, blog.Body, AnyTime{}, AnyTime{}).
// 		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(blog.ID))
//
// 	if _, err = rw.Create(blog); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
//
// 	// 使用されたモックDBが期待通りの値を持っているかを検証
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
//
// func TestRw_unHappyCreate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//
// 	rw := blogRW.New(db)
// 	blog := testData.Blog()
//
// 	mock.ExpectQuery(regexp.QuoteMeta(
// 		`INSERT INTO blogs (title, body, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING id`)).
// 		WillReturnError(fmt.Errorf("some error"))
//
// 	if _, err = rw.Create(blog); err == nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
//
// 	// err == nil -> つまり、エラーがあればpassするということ
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
//
// func TestRw_happyUpdate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//
// 	rw := blogRW.New(db)
// 	blog := testData.Blog()
//
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE blogs SET title = $2, body = $3, updated_at = $4 WHERE id = $1`)).
// 		WithArgs(blog.ID.Value, blog.Title, blog.Body, AnyTime{}).
// 		WillReturnResult(sqlmock.NewResult(1, 5))
//
// 	// モック化されたDBを用いてテスト対象関数を実行
// 	if _, err = rw.Update(blog.ID.Value(), blog); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
//
// 	// 使用されたモックDBが期待通りの値を持っているかを検証
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
//
// func TestRw_unHappyUpdate(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//
// 	rw := blogRW.New(db)
// 	blog := testData.Blog()
//
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE blogs SET title = $2, body = $3, updated_at = $4 WHERE id = $1`)).
// 		WithArgs(blog.ID, blog.Title, blog.Body, AnyTime{}).
// 		WillReturnError(fmt.Errorf("some error"))
//
// 	// モック化されたDBを用いてテスト対象関数を実行
// 	if _, err = rw.Update(blog.ID.Value(), blog); err == nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
//
// 	// 使用されたモックDBが期待通りの値を持っているかを検証
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
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
