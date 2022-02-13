package blogDao_test

import (
	"clean_architecture/golang/adapters/dao/blogDao"
	"clean_architecture/golang/testData"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTest(name string) (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"test-db",
		5432,
		"postgresql",
		"postgresql",
		"test-api",
	)

	fmt.Println("sentinel1")
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		fmt.Println("sentinel2")
		fmt.Println("not connected")
	}

	txdb.Register(name, "postgres", conn)
	fmt.Println("sentinel3")

	return db, nil
}

func prepare(name string, seeds []interface{}) (*gorm.DB, error) {
	fmt.Println("------------")
	fmt.Println(name)
	fmt.Println(seeds)
	db, err := NewTest(name)
	if err != nil {
		fmt.Println("sentinel4")
		return nil, err
	}
	fmt.Println("hoge")
	for _, s := range seeds {
		fmt.Println("hoge1")
		fmt.Println(s)
		if err := db.Create(s).Error; err != nil {
			fmt.Println("sentinel5")
			return nil, err
		}
		fmt.Println("hoge2")
	}
	fmt.Println("hoge3")
	return db, nil
}

func TestGetById_Success(t *testing.T) {
	t.Parallel()

	fmt.Println("sentinel11")
	blog := testData.Blog()
	fmt.Println("sentinel12")
	// 変数ではなくアドレス入れる必要あり
	seeds := []interface{}{
		&blogDao.BlogDto{
			Title: blog.Title.Value,
			Body:  blog.Body.Value,
		},
	}
	// fmt.Println(db)
	// fmt.Println(err)
	db, err := prepare("users_dao", seeds)
	if err != nil {
		fmt.Println("sentinel7")
		t.Fatal(err)
	}

	fmt.Println("sentinel8")
	rw := blogDao.New(db)
	// d := rw.GetById()

	tests := []struct {
		name   string
		give   int
		wantID int
		err    bool
	}{
		{
			name:   "success",
			give:   1,
			wantID: 1,
		},
		{
			name: "not found",
			give: 2,
			err:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fmt.Println("haitta inside")
			// rw.Create(testData.NewBlog())

			got, aerr := rw.GetById(tt.give)
			if tt.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
				assert.Equal(t, tt.wantID, got.ID.Value)
			}

		})
	}

}
