// Package command 命令模式
// Blog: https://lailin.xyz/post/command.html
// 这是示例一，采用将函数封装为对象的方式实现，
// 示例说明:
// 假设现在有一个游戏服务，我们正在实现一个游戏后端
// 使用一个 goroutine 不断接收来自客户端请求的命令，并且将它放置到一个队列当中
// 然后我们在另外一个 goroutine 中来执行它
package command

import "fmt"

// ICommand 命令
type ICommand interface {
	Execute() error
}

// StartCommand 游戏开始运行
type StartCommand struct{}

// NewStartCommand NewStartCommand
func NewStartCommand( /*正常情况下这里会有一些参数*/ ) *StartCommand {
	return &StartCommand{}
}

// Execute Execute
func (c *StartCommand) Execute() error {
	fmt.Println("game start")
	return nil
}

// ArchiveCommand 游戏存档
type ArchiveCommand struct{}

// NewArchiveCommand NewArchiveCommand
func NewArchiveCommand( /*正常情况下这里会有一些参数*/ ) *ArchiveCommand {
	return &ArchiveCommand{}
}

// Execute Execute
func (c *ArchiveCommand) Execute() error {
	fmt.Println("game archive")
	return nil
}
