package blogDao_test

import (
	"clean_architecture/golang/adapters/dao/blogDao"
	"clean_architecture/golang/testData"
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// AnyTimeはテストを通すための設定。公式を参考に実装
type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestRw_GetAll_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rw := blogDao.New(db)
	blog := testData.Blog()

	// DBモック用意
	mock.ExpectQuery(regexp.QuoteMeta(blogDao.GetAllSql)).
		WithArgs().
		WillReturnRows(mock.NewRows([]string{
			"id",
			"title",
			"body",
			"created_at",
			"updated_at",
		}).AddRow(
			blog.ID.Value,
			blog.Title.Value,
			blog.Body.Value,
			blog.CreatedAt,
			blog.UpdatedAt,
		))

	// モック化されたDBを用いてテスト対象関数を実行
	blogs, err := rw.GetAll()

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	assert.NotEmpty(t, blogs)
	assert.NoError(t, err)
	assert.Equal(t, blogs.Size(), 1)

	//// 使用されたモックDBが期待通りの値を持っているかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRw_GetAll_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rw := blogDao.New(db)
	blog := testData.Blog()

	// DBモック用意
	mock.ExpectQuery(regexp.QuoteMeta(blogDao.GetAllSql)).
		WithArgs().
		WillReturnRows(mock.NewRows([]string{
			"id",
			"title",
			"body",
			"created_at",
			"updated_at",
		}).AddRow(
			blog.ID.Value,
			blog.Title.Value,
			blog.Body.Value,
			blog.CreatedAt,
			blog.UpdatedAt,
		)).
		WillReturnError(fmt.Errorf("some error"))

	_, err = rw.GetAll()

	// モック化されたDBを用いてテスト対象関数を実行
	if err == nil {
		t.Errorf("Expect error happens:%s", err)
	}

	// 使用されたモックDBが期待通りの値を持っているかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
