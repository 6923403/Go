package main

import "fmt"

func main() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf(" a = %d\n", a )
	fmt.Printf(" *ptr = %d\n", ptr )
	fmt.Printf(" **pptr = %d\n", **pptr)

	/* address*/
	fmt.Printf("\n a = %d\n", &a )
	fmt.Printf(" *ptr = %d\n", &ptr )
	fmt.Printf(" **pptr = %d\n", &pptr)
}