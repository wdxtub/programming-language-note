package main

import "fmt"

// 声明一个函数类型，用做回调，这里用于输出统一的日志信息
type callback func(string) string

func main() {
	// 传入函数作为回调
	testCallback("hello", callbackSample)
	// 使用匿名函数作为回调
	testCallback("world", func(str string) string {
		fmt.Println("[CALLBACK]", str)
		return "[CALLBACK]" + str
	})
}

func testCallback(str string, f callback) {
	f(str)
}

// 这个函数需要和之前声明的一样
func callbackSample(str string) string {
	fmt.Println("[CALLBACK]", str)
	return "[CALLBACK]" + str
}
