package main

import "fmt"

func main() {
	// 声明一个变量并初始化
	var sampleInitStr = "wdxtub.com"
	fmt.Println(sampleInitStr)

	// 字符串
	var sampleStr string
	fmt.Println(sampleStr)

	// 整型、浮点型初始化
	var sampleInt int
	fmt.Println(sampleInt)
	var sampleFloat float32
	fmt.Println(sampleFloat)

	// 布尔值初始化
	var sampleBool bool
	fmt.Println(sampleBool)

	var t1 *int
	fmt.Println(t1)
	var t2 []int
	fmt.Println(t2)
	var t3 map[string]int
	fmt.Println(t3)
	var t4 chan int
	fmt.Println(t4)
	var t5 func(string) int
	fmt.Println(t5)
	var t6 error // error 是接口
	fmt.Println(t6)
}
