package main // we are using package main which is already defined by GO-runtime

import (
	"fmt"
	"reflect"
)

func main() {
	//println("Hello World!!") //here(println/print) it's a build-in function

	var age uint8 = 38

	fmt.Println("Age:", age, "Type of age:", reflect.TypeOf(age))
	fmt.Printf("\nAge: %v Type of age:  %T\n", age, age)

	var num1 = 100
	var num2 = 25.34

	num3 := 1239 // this is done during the compile time

	fmt.Println(num1, num2, num3, reflect.TypeOf(num1), reflect.TypeOf(num2), reflect.TypeOf(num3))

	//type inference
	var (
		a int
		b float32
		c string
		d bool
	)
	fmt.Println(a, b, c, d)
}
