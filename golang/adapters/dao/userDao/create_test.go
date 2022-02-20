package userDao_test

import (
	"clean_architecture/golang/adapters/dao/userDao"
	"clean_architecture/golang/domains"
	"clean_architecture/golang/testData"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	user := testData.User()
	seeds := []interface{}{}

	db, err := Prepare("user_create_dao", seeds)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatal(err)
	}

	rw := userDao.New(db)

	tests := []struct {
		name      string
		user      *domains.User
		wantError error
	}{
		{
			name:      "Create a user",
			wantError: nil,
		},
		// {
		// 	name:      "Return not found error",
		// 	user:      &dummyUser,
		// 	wantError: gorm.ErrRecordNotFound,
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := rw.Create(user)
			if tt.wantError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tt.wantError)
			}
		})
	}
}
