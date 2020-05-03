package main

import "fmt"

func main(){
	te3()
}

func te3(){
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Printf("a = %T \n", a)
	fmt.Printf("b = %T \n", b)
	fmt.Printf("c = %T \n", c)

	ptr = &a
	fmt.Printf("a = %d \n", a)
	fmt.Printf("ptr = %d \n", *ptr)
}

func te2(){
	var a uint = 60
	var b uint = 13
	var c uint = 0

	c = a & b
	fmt.Printf("c(a&b) = %d \n", c)

	c = a | b
	fmt.Printf("c(a|b) = %d \n", c)

	c = a ^ b
	fmt.Printf("c(a^b) = %d \n", c)

	c = a << 2
	fmt.Printf("c(a<<2) = %d \n", c)

	c = a >> 2
	fmt.Printf("c(a>>2) = %d \n", c)
}

func te1(){
	var a bool = true
	var b bool = false
	if(a && b){
		fmt.Println("first line is true \n")
	}

	if(a || b){
		fmt.Println("second line is true \n")
	}
	a = false
	b = true

	if(a && b){
		fmt.Printf("three line is true \n")
	}	else{
		fmt.Printf("three line is false \n")
	}

	if(!(a && b)){
		fmt.Printf("four line is true")
	}
}
