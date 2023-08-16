package main

import "fmt"

// defer --> in the sense it makes the fn to execute at the end of the scope
func main() {
	defer func() {
		defer fmt.Println("Hi this is first") // this will be executed at the end of the scope
		fmt.Println("hiiiiiiii")
	}() // this is executed at the last
	defer fmt.Println("where is this printed ?") // this will be printed before the first defer
	//this is because it uses LIFO method or stack
	fmt.Println("This is the last")
}
