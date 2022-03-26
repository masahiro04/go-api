package user_test

import (
	"go-api/domains/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNameSuccess(t *testing.T) {
	input := "hogehoge@example.com"
	newName, err := user.NewName(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newName.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewNameFailWithBlankString(t *testing.T) {
	input := ""
	newName, err := user.NewName(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newName.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdateNameSuccess(t *testing.T) {
	input := "hogehoge@example.com"
	updatedName, err := user.UpdateName(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedName.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateNameFail(t *testing.T) {
	input := ""
	updatedName, err := user.UpdateName(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedName.Value, input)
		assert.NotNil(t, err)
	})
}

func TestNameValueSuccess(t *testing.T) {
	input := "hogehoge@example.com"
	Name, _ := user.NewName(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, Name.Value, input)
	})
}
