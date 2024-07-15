package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@example.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
}

func TestUserComparePassword(t *testing.T) {
	user, _ := NewUser("John Doe", "john@example.com", "123456")
	assert.NotEqual(t, "123456", user.Password)
	assert.Nil(t, user.ComparePassword("123456"))   // correct password
	assert.NotNil(t, user.ComparePassword("12345")) // incorrect password
}
