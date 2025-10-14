package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 1) // 有缓冲通道，容量为1
	signals := make(chan bool)

	// 启动一个goroutine模拟后台任务
	go func() {
		time.Sleep(1 * time.Second)
		messages <- "Hello from goroutine"
	}()

	// 主程序继续执行其他任务
	fmt.Println("Main thread working...")
	time.Sleep(500 * time.Millisecond)

	// 非阻塞检查是否有消息到达
	select {
	case msg := <-messages:
		fmt.Println("Early message:", msg)
	default:
		fmt.Println("No early message yet")
	}

	// 等待一会儿再检查
	time.Sleep(600 * time.Millisecond)
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity detected")
	}

	// 模拟信号发送
	go func() {
		time.Sleep(100 * time.Millisecond)
		signals <- true
	}()

	// 带超时的选择
	select {
	case msg := <-messages:
		fmt.Println("Final message:", msg)
	case sig := <-signals:
		fmt.Println("Final signal:", sig) // 可能会执行这里
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation timed out")
	}
}
