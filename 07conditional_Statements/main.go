package main

import "fmt"

func main() {
	//isEven(101)

	isValid(100)
}

func isEven(num int) {
	if num%2 == 0 {
		fmt.Println("It is even")
	} else {
		fmt.Println("It's an odd number")
	}
}

func isValid(num int) {
	if num < 50 {
		fmt.Println(num + 1)
	} else if num > 50 && num < 100 {
		fmt.Println(num - 1)
	} else if num > 100 {
		fmt.Println(num - num)
	} else {
		fmt.Println(num)
	}
}
