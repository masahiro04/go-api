package models_test

import (
	"go-api/domains/models"
	"go-api/test/factories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsers_Success(t *testing.T) {
	_users := factories.Users(5)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, len(_users.Value), 4)
	})
}

func TestUsersSize_Success(t *testing.T) {
	_users := factories.Users(5)

	t.Run("", func(t *testing.T) {
		assert.Equal(t, _users.Size(), 4)
	})
}

func TestUsersEmptyUsers_Success(t *testing.T) {
	newUsers := models.EmptyUsers()

	t.Run("", func(t *testing.T) {
		assert.Equal(t, newUsers.Size(), 0)
	})
}

func TestUsersApplyLimitAndOffset_Success(t *testing.T) {
	var _users = factories.Users(5)

	t.Run("complete", func(t *testing.T) {
		assert.Equal(t, _users.Value, _users.ApplyLimitAndOffset(100, 0))
		assert.Equal(t, _users.Value, _users.ApplyLimitAndOffset(4, 0))
		assert.Equal(t, _users.Value, _users.ApplyLimitAndOffset(4, -1))
	})
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, models.EmptyUsers().Value, _users.ApplyLimitAndOffset(100, 10))
		assert.Equal(t, models.EmptyUsers().Value, _users.ApplyLimitAndOffset(3, 4))
		assert.Equal(t, models.EmptyUsers().Value, _users.ApplyLimitAndOffset(-1, 0))
	})
}
