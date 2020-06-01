package main

import "fmt"

func te(){
	var grade string = "B"
	var marks int = 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50, 60, 70: grade = "C"
		default: grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("Go \n")
	case grade == "B":
		fmt.Printf("Goo \n")
	case grade == "D":
		fmt.Printf("Gooo \n")
	case grade == "F":
		fmt.Printf("Goooo \n")
	default:
		fmt.Printf("Null \n")
	}
	fmt.Printf("You is %s", grade);
}

func te2() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x type is : %T", i)
	case int:
		fmt.Printf("x type is int. \n")
	case float64:
		fmt.Printf("x type is float64. \n")
	case func(int) float64:
		fmt.Printf("x type is func(int). \n")
	case bool, string:
		fmt.Printf("x is bool or string. \n")
	default:
		fmt.Printf("Null")
	}
}

func main() {
	switch {
	case false:
		fmt.Println("1. case is false \n")
		fallthrough /* 不判断下一条语句 */
	case true:
		fmt.Println("2. case is true \n")
		fallthrough
	case false:
		fmt.Println("3. case is false \n")
		fallthrough
	case true:
		fmt.Println("4. case is true \n")
	case false:
		fmt.Println("5. case is false \n")
		fallthrough
	default:
		fmt.Println("6. default case \n")
	}
}
