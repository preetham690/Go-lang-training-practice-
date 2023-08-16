package main

import "fmt"

func main() {

	a := 10
	b := 20

	defer fmt.Println(a, b, add(a, b)) // it behaves diff, when we use inline functions
	//because it maintains it's own runstack
	//when we use inline fn then the value will be passed to the fn but waits till the end

	defer func() {
		c := a + b
		fmt.Println(a, b, c)
	}()

	a = 100
	b = 200

	fmt.Println(a, b, add(a, b))
}

func add(a, b int) int {
	return a + b
}
