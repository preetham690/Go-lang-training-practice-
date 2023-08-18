package main

import (
	"encoding/json"
	"fmt"
)

// command code --> go run -gcflags="m" main.go
func main() {
	var i1 int = 100
	var ptr1 *int = &i1

	arr := [5]int{1, 2, 4, 5, 6}
	slice := []int{1, 2, 3, 4}

	println(slice)

	fmt.Println(arr)
	fmt.Println(i1, ptr1)

	println(sump(2, 3))

	s1 := sample{num: "01", name: "pree"}
	bytes, _ := json.Marshal(s1)
	fmt.Println(string(bytes))

	//println(*sump(10, *sump(20, 30)))
}

type sample struct {
	num, name string
}

func sump(a, b int) *int {
	c := a + b
	return &c // here for sure it will be stored to heap, because we are sending the pointer which has the scope of that fucntion
}
