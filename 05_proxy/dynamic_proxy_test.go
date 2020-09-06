package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_generate(t *testing.T) {
	want := `package proxy

type UserProxy struct {
	child *User
}

func NewUserProxy(child *User) *UserProxy {
	return &UserProxy{child: child}
}

func (p *UserProxy) Login(username, password string) (r0 error) {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	r0 = p.child.Login(username, password)

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return r0
}
`
	got, err := generate("./static_proxy.go")
	require.Nil(t, err)
	assert.Equal(t, want, got)
}
