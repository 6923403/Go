package main

import "fmt"
/*
* 以下描述了 select 语句的语法：
每个 case 都必须是一个通信
所有 channel 表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行，其他被忽略。
如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
否则：
如果有 default 子句，则执行该语句。
如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
 */

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
		case i1 = <-c1:
			fmt.Printf("received ", i1, "form c1 \n")
		case c2 <- i2:
			fmt.Printf("sent ", i2, "to c2 \n")
		case i3, ok := (<-c3):
			if ok {
				fmt.Printf("received ", i3, "from c3 \n")
			} else {
				fmt.Printf("c3 is closed \n")
			}
	default:
		fmt.Printf("no communication \n")
	}
}
