package main

import "fmt"

func main() {
	//here there's a cascading effect
	func() {
		func() {
			r := Divide(100, 0) // this creates a panic -> which in turn crashes the above fn and above and finally main will crash
			fmt.Println(r)
		}()
		fmt.Println("hi")
	}()
	fmt.Println("hello")
}

func Divide(n, div int) int {
	return n / div
}
