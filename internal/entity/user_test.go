package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("giovane", "giovane@email.com", "1234")
	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, user, "User should not be nil")
	assert.Equal(t, "giovane", user.Name, "Name should be giovane")
	assert.Equal(t, "giovane@email.com", user.Email, "Email should be giovane@email.com")
	assert.NotEmpty(t, user.Password, "Password should not be empty")
	assert.True(t, user.ValidatePassword("1234"), "Password should be hashed and valid")
}
