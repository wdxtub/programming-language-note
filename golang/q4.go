package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("1 - false")
		fallthrough
	case true:
		fmt.Println("2 - true")
		fallthrough
	case false:
		fmt.Println("3 - false")
		fallthrough
	case true:
		fmt.Println("4 - true")
	case false:
		fmt.Println("5 - false")
		fallthrough
	default:
		fmt.Println("6 - default")
	}

}
