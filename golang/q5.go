package main

import (
	"fmt"
	"time"
)

// Sender 用来发送信号
func Sender(ch chan int, stopCh chan bool) {
	i := 220
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

func main() {
	// 创建一个发送单个 int 的 channel
	ch := make(chan int)
	var c int
	// 创建一个发送单个 bool 的 channel，用于停止程序
	stopCh := make(chan bool)

	go Sender(ch, stopCh)

	for {
		select {
		case c = <-ch: // 从 ch 中读取
			fmt.Println("Receiver 1", c)
		case s := <-ch: // 从 ch 中读取
			fmt.Println("Receiver 2", s)
		case t := <-ch: // 从 ch 中读取
			fmt.Println("Receiver 3", t)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}
