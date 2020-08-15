/*
# 1
var identifier []type

# 2
var slice1 []type = make([]type, len)
也可以简写为
slice1 := make([]type, len)

#3
make([]T, length, capacity)
 */

package main

import "fmt"

func main() {
	//var numbers = make([]int,3,5)
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	pslice(numbers2)
	numbers2 = apendslice(numbers2)

	//same len and capcity
	numbers3 := make([]int, len(numbers2), cap(numbers2))
	copy(numbers3,numbers2)

	pslice(numbers2)
	pslice(numbers3)
}

func apendslice(x []int) []int {
	x = append(x, 10)
	return x
}

func pslice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}