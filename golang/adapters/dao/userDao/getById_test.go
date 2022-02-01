package userDao_test

import (
	"clean_architecture/golang/adapters/dao/userDao"
	"clean_architecture/golang/testData"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRw_happyGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rw := userDao.New(db)
	user := testData.User()

	// DBモック用意
	mock.ExpectQuery(regexp.QuoteMeta(userDao.GetByIdSql)).
		WithArgs().
		WillReturnRows(mock.NewRows([]string{
			"id",
			"name",
			"email",
			"created_at",
			"updated_at",
		}).AddRow(
			user.ID.Value,
			user.Name.Value,
			user.Email.Value,
			user.CreatedAt,
			user.UpdatedAt,
		))

	// モック化されたDBを用いてテスト対象関数を実行
	if _, err := rw.GetById(user.ID.Value); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// 使用されたモックDBが期待通りの値を持っているかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRw_unHappyGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rw := userDao.New(db)
	user := testData.User()

	// DBモック用意
	mock.ExpectQuery(regexp.QuoteMeta(userDao.GetByIdSql)).
		WithArgs().
		WillReturnError(fmt.Errorf("some error"))

	// モック化されたDBを用いてテスト対象関数を実行
	if _, err := rw.GetById(user.ID.Value); err == nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// 使用されたモックDBが期待通りの値を持っているかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
