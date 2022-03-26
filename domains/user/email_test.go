package user_test

import (
	"go-api/domains/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmailSuccess(t *testing.T) {
	input := "hogehoge@example.com"
	newEmail, err := user.NewEmail(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newEmail.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestNewEmailFailWithBlankString(t *testing.T) {
	input := ""
	newEmail, err := user.NewEmail(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newEmail.Value, input)
		assert.NotNil(t, err)
	})
}

func TestNewEmailFailWithNotCorrectFormat(t *testing.T) {
	input := "hogehoge"
	newEmail, err := user.NewEmail(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, newEmail.Value, input)
		assert.NotNil(t, err)
	})
}

func TestUpdateEmailSuccess(t *testing.T) {
	input := "hogehoge@example.com"
	updatedEmail, err := user.UpdateEmail(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedEmail.Value, input)
		assert.Equal(t, err, nil)
	})
}

func TestUpdateEmailFail(t *testing.T) {
	input := ""
	updatedEmail, err := user.UpdateEmail(&input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, updatedEmail.Value, input)
		assert.NotNil(t, err)
	})
}

func TestEmailValueSuccess(t *testing.T) {
	input := "hogehoge@example.com"
	Email, _ := user.NewEmail(input)
	t.Run("hoge", func(t *testing.T) {
		assert.Equal(t, Email.Value, input)
	})
}
