package main

import (
	"fmt"
	"unsafe"
)

func main() {
	t4()
}

func t4() {
	const (
		i = 1 << iota //iota = 0
		j = 3 << iota //iota = 1
		k
		l
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)
}

func t3() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)
}

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func t2() {
	println(a, b, c)
}

func t1() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d", area)
	println()
	println(a, b, c)
}
