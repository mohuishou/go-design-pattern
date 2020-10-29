// Package command 命令模式
// Blog: https://lailin.xyz/post/command.html
// 这是示例二，采用将直接返回一个函数，不用对象
// 示例说明:
// 假设现在有一个游戏服务，我们正在实现一个游戏后端
// 使用一个 goroutine 不断接收来自客户端请求的命令，并且将它放置到一个队列当中
// 然后我们在另外一个 goroutine 中来执行它
package command

import "fmt"

// Command 命令
type Command func() error

// StartCommandFunc 返回一个 Command 命令
// 是因为正常情况下不会是这么简单的函数
// 一般都会有一些参数
func StartCommandFunc() Command {
	return func() error {
		fmt.Println("game start")
		return nil
	}
}

// ArchiveCommandFunc ArchiveCommandFunc
func ArchiveCommandFunc() Command {
	return func() error {
		fmt.Println("game archive")
		return nil
	}
}
