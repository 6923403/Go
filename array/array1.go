package main

import "fmt"

func main() {
	var n[10] int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("n[%d] = %d \n", j, n[j])
	}

	var ccc float32
	ccc = getAverage(n, 10);
	fmt.Printf("\n ccc = %f\n", ccc)
}

func getAverage(arr [10]int, size int) float32 {
	var i int
	var sum float32
	var ccc float32

	for i = 0; i < size; i++ {
		sum += float32(arr[i])
	}

	ccc = sum / float32(size)

	return ccc;
}