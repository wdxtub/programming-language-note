package main

import "fmt"

func channelFibo(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n ; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 这里一定要关闭，不然会阻塞
}

func main() {
	c := make(chan int, 15)
	go channelFibo(cap(c), c)

	for result := range c {
		fmt.Println(result)
	}
}