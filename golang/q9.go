package main

import (
	"fmt"
	"runtime"
	"time"
)

// 简易计数器
func simpleCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 第一题
func question1() {
	counter := simpleCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	counter = simpleCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		fmt.Printf("%p\n", &base) // 这里可以看到都是同一个地址
		base += i
		return base
	}

	sub := func(i int) int {
		fmt.Printf("%p\n", &base) // 这里可以看到都是同一个地址
		base -= i
		return base
	}
	return add, sub
}

// 第二题
func question2() {
	f1, f2 := calc(100)

	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))

	// 实际上是 100+1-2+3-4+5-6
}

// 第三题
func question3() {
	runtime.GOMAXPROCS(1) // 设置最大可同时使用的 CPU 核数为 1

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("3.1:", i)
		}()
	}
	time.Sleep(3 * time.Second)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("3.2:", i)
		}(i)
	}
	time.Sleep(3 * time.Second)
}

func main() {
	// 总共有多道题目，每个题目封装到一个函数中
	question1()
	question2()
	question3()
}
