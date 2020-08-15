package main

import "fmt"

const Max  = 5

func main() {
	test2()
}

func test2() {
	a := []int{1, 2, 3, 4, 5}
	var i, j int
	var ptr [Max] *int

	for i = 0; i < Max; i++ {
		ptr[i] = &a[i]
	}

	for j = 0; j < Max; j++ {
		fmt.Printf("a[%d] = %d\n", j, *ptr[j])
	}
}

func test1() {
	a := []int{1, 2, 3}
	var i int

	for i = 0; i < Max; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}