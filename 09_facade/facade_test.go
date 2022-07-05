package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Login(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.User.Login(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.LoginOrRegister(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestUserService_Register(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.User.Register(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test register"}, user)
}
