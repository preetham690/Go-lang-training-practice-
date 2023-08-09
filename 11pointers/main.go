package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int = 100
	var ptr *int = &num // here * is nothing but the value in that particular address

	var prt1 **int = &ptr // it points to the address of another variable
	fmt.Println("Value of num is: ", *ptr, "\nAddress of num is: ", &ptr)

	fmt.Println("Value of num is:", **prt1, "\nAddress of num is:", prt1)

	temp := 10
	var prt2 *int = nil

	//better to check whether the pointer is nil are not
	increment(prt2)
	fmt.Println(temp)
}

// incrementing variable using the pointer concept
func increment(num *int) error {
	if num == nil {
		err := errors.New("the pointer is nil")
		fmt.Println(err)
		return err
	}

	*num++

	return nil
}
