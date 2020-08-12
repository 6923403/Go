package main

import "fmt"

func main() {
	var a int = 5
	var b int = 10
	var ret int

	ret = max(a, b)
	fmt.Printf("max value = %d \n", ret)
}
/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
/* 声明局部变量 */
var result int

if (num1 > num2) {
	result = num1
} else {
	result = num2
}
	return result
}