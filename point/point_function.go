package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("a = %d, b = %d \n", a, b)
	swap(&a, &b)
	fmt.Printf("swap: a = %d, b = %d", a, b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}