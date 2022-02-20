package userDao_test

import (
	"clean_architecture/golang/adapters/dao/userDao"
	"clean_architecture/golang/domains"
	"clean_architecture/golang/testData"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetById(t *testing.T) {

	user := testData.User()
	dummyUser := testData.UserWithID(100) // 100でnot foud起こす
	seeds := []interface{}{
		&userDao.UserDto{
			ID:    user.ID.Value,
			Name:  user.Name.Value,
			Email: user.Email.Value,
		},
	}

	db, err := Prepare("user_dao", seeds)

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
			name:      "Get a user",
			user:      &user,
			wantError: nil,
		},
		{
			name:      "Return not found error",
			user:      &dummyUser,
			wantError: gorm.ErrRecordNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := rw.GetById(tt.user.ID.Value)
			if tt.wantError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err, tt.wantError)
			}
		})
	}
}
