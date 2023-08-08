package main

import "fmt"

var (
	n1 int
	n2 float32
	n3 float64
)

func main() {
	var str1 string = "Hello world!!"

	//strings are immutable in go from the runtime perspective
	//in the sense we can not reassign diff value for the same variable -> in the bg
	//it will create a anothe space for that variable and assign the value , it won't clear the previous location space it will be still held in the same memory
	str1 = "Hi vs"
	fmt.Printf("Lenght of string 1 is: %v and the string is %v", len(str1), str1)
}
