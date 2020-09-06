package proxy

import (
	"log"
	"time"
)

// IUser IUser
type IUser interface {
	Login(username, password string) error
}

// User 用户
// @proxy IUser
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// 不实现细节
	return nil
}

// UserProxy 代理类
type UserProxy struct {
	user *User
}

// NewUserProxy NewUserProxy
func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

// Login 登录，和 user 实现相同的接口
func (p *UserProxy) Login(username, password string) error {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	// 这里是原有的业务逻辑
	if err := p.user.Login(username, password); err != nil {
		return err
	}

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}
