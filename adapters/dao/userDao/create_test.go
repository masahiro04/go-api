package userDao_test

import (
	"go-api/adapters/dao/userDao"
	"go-api/domains"
	factories "go-api/test/factories"
	testhelpers "go-api/test/testHelpers"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	user := factories.User()
	seeds := []interface{}{}

	db, err := testhelpers.Prepare("user_create_dao", seeds)

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
