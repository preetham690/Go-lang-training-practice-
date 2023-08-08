package main

import (
	"fmt"
	"reflect"
)

const PI = 3.14

const PI_SQR = PI * PI

type integer = int //typeDef

func main() {
	fmt.Println(reflect.TypeOf(PI_SQR))

	var a int = 2

	fmt.Println(reflect.TypeOf(a))

	var num1 integer = 100001234
	var num2 rune = 'a' //nothing but the int32
	var num3 int32 = 'b'
	var num4 uint8 = 'A'

	var num5 rune = 'ಕ' //here it will print the Unicode of that particular character
	//this is how they map the characters in translator "using UNIcode"

	fmt.Println(num1, num2, num3, num4, num5)
}
