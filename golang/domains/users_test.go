package domains_test

import (
	"clean_architecture/golang/domains"
	"clean_architecture/golang/testData"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsersSuccess(t *testing.T) {
	_users := testData.Users(5)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, len(_users.Value), 4)
	})
}

func TestUsersSizeSuccess(t *testing.T) {
	_users := testData.Users(5)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _users.Size(), 4)
	})
}

func TestUsersEmptyUsersSuccess(t *testing.T) {
	newUsers := domains.EmptyUsers()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newUsers.Size(), 0)
	})
}

func TestUsersApplyLimitAndOffset(t *testing.T) {
	var _users = testData.Users(5)

	t.Run("complete", func(t *testing.T) {
		assert.Equal(t, _users.Value, _users.ApplyLimitAndOffset(100, 0))
		assert.Equal(t, _users.Value, _users.ApplyLimitAndOffset(4, 0))
		assert.Equal(t, _users.Value, _users.ApplyLimitAndOffset(4, -1))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, domains.EmptyUsers().Value, _users.ApplyLimitAndOffset(100, 10))
		assert.Equal(t, domains.EmptyUsers().Value, _users.ApplyLimitAndOffset(3, 4))
		assert.Equal(t, domains.EmptyUsers().Value, _users.ApplyLimitAndOffset(-1, 0))
	})
}
