package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "wdxtub"
		e
		f
		g = 10086
		h
		i
		j = iota
		k
		l
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)

	const (
		z = 1 << iota
		y
		x
		w = 2 << iota
		v
		u
	)
	fmt.Println(z, y, x, w, v, u)
}
