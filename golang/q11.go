package main

import "fmt"

func main() {
	oneInt := 17
	oneFloat := 17.2

	int2Float := float32(oneInt)
	float2Int := int(oneFloat)

	fmt.Println(int2Float)
	fmt.Println(float2Int)
	oneFloat = 17.6
	float2Int = int(oneFloat)
	fmt.Println(float2Int)
}
