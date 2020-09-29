package main

import (
	"fmt"
)

// 第一种交换方式
func swap1(x, y int) {
	x, y = y, x
	fmt.Println("swap1 函数交换", x, y)
}

// 第二种交换方式
func swap2(x, y *int) {
	*x, *y = *y, *x
	fmt.Println("swap1 函数交换", *x, *y)
}

// 修改 map
func modifyMap(item map[string]int) {
	item["a"] = 314
	fmt.Println("modifyMap 函数修改", item)
}

// 修改 切片
func modifySlice(item []int) {
	item[0] = 314
	fmt.Println("modifySlice 函数修改", item)
}

// 修改 数组
func modifyArray(item [3]int) {
	item[0] = 314
	fmt.Println("modifyArray 函数修改", item)
}

// 修改数组指针
func modifyArrayPtr(item *[3]int) {
	item[0] = 314
	fmt.Println("modifyArrayPtr 函数修改", *item)
}

func main() {
	x := 314
	y := 220
	fmt.Println("x y 交换前", x, y)
	swap1(x, y)
	fmt.Println("x y swap1 交换后", x, y)
	swap2(&x, &y)
	fmt.Println("x y swap2 交换后", x, y)

	arraySample := [3]int{1, 2, 3}
	fmt.Println("arraySample 修改前", arraySample)
	modifyArray(arraySample)
	fmt.Println("arraySample 修改后", arraySample)
	modifyArrayPtr(&arraySample)
	fmt.Println("modifyArrayPtr 修改后", arraySample)

	sliceSample := []int{1, 2, 3}
	fmt.Println("sliceSample 修改前", sliceSample)
	modifySlice(sliceSample)
	fmt.Println("sliceSample 修改后", sliceSample)

	mapSample := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("mapSample 修改前", mapSample)
	modifyMap(mapSample)
	fmt.Println("mapSample 修改后", mapSample)
}
