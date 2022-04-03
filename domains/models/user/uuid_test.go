package user_test

import (
	"go-api/domains/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID_Success(t *testing.T) {
	input := "hogehoge@example.com"
	newUUID, err := user.NewUUID(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newUUID.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewUUIDWithBlankString_Fail(t *testing.T) {
	input := ""
	newUUID, err := user.NewUUID(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newUUID.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdateUUID_Success(t *testing.T) {
	input := "hogehoge@example.com"
	updatedUUID, err := user.UpdateUUID(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedUUID.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateUUID_Fail(t *testing.T) {
	input := ""
	updatedUUID, err := user.UpdateUUID(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedUUID.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUUIDValue_Success(t *testing.T) {
	input := "hogehoge@example.com"
	UUID, _ := user.NewUUID(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, UUID.Value, input)
	})
}
