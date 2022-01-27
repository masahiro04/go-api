package blogDao_test

import (
	"clean_architecture/golang/adapters/dao/blogDao"
	"clean_architecture/golang/testData"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRw_happyCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rw := blogDao.New(db)
	blog := testData.Blog()
	//id := 1
	mock.ExpectQuery(regexp.QuoteMeta(blogDao.CreateSql)).
		WithArgs(blog.Title.Value, blog.Body.Value, AnyTime{}, AnyTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(blog.ID.Value))

	if _, err = rw.Create(blog); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// 使用されたモックDBが期待通りの値を持っているかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRw_unHappyCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rw := blogDao.New(db)
	blog := testData.Blog()

	mock.ExpectQuery(regexp.QuoteMeta(blogDao.CreateSql)).
		WillReturnError(fmt.Errorf("some error"))

	if _, err = rw.Create(blog); err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// err == nil -> つまり、エラーがあればpassするということ
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
