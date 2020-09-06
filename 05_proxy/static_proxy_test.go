package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserProxy_Login(t *testing.T) {
	proxy := NewUserProxy(&User{})

	err := proxy.Login("test", "password")

	require.Nil(t, err)
}
