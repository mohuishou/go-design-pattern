package command

import (
	"fmt"
	"testing"
	"time"
)

func TestDemoFunc(t *testing.T) {
	// 用于测试，模拟来自客户端的事件
	eventChan := make(chan string)
	go func() {
		events := []string{"start", "archive", "start", "archive", "start", "start"}
		for _, e := range events {
			eventChan <- e
		}

	}()
	defer close(eventChan)

	// 使用命令队列缓存命令
	commands := make(chan Command, 1000)
	defer close(commands)

	go func() {
		for {
			// 从请求或者其他地方获取相关事件参数
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command Command
			switch event {
			case "start":
				command = StartCommandFunc()
			case "archive":
				command = ArchiveCommandFunc()
			}

			// 将命令入队
			commands <- command
		}
	}()

	for {
		select {
		case c := <-commands:
			c()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
