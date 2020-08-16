package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type APhone struct {

}

func (AP APhone) call() {
	fmt.Println("I am AP, I can call you!")
}

type IPhone struct {

}

func (IP IPhone) call() {
	fmt.Println("I am IP, I can call you!")
}

func main() {
	var phone Phone

	phone = new(APhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}