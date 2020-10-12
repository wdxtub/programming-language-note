package main

import (
	"fmt"
)

func gossip(s string) {
	fmt.Println("[Gossip]", s)
}

func main() {
	go gossip("wdxtub")

	// 如果想要看到结果，需要让主进程等待
	// time.Sleep(100 * time.Millisecond)
}
