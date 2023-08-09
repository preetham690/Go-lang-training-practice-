package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{11, 22, 33, 44}

	fmt.Println(sum(arr[:]...))
	fmt.Println(sum(slice...))
}

//variadic parameter must be the last parameter
//also there is only one variadic func for a program

func sum(args ...int) int {
	_sum := 0

	for _, v := range args {
		_sum += v
	}

	return _sum
}
