// 常规的通过通道发送和接收数据是阻塞的。然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。

package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select { // 这里是一个非阻塞接收的例子。如果在 messages 中存在，然后 select 将这个值带入 <-messages case中。如果不是，就直接到 default 分支中。
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	time.Sleep(time.Second * 2)

	select { // 我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。这里我们试图在 messages和 signals 上同时使用非阻塞的接受操作。
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
