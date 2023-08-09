/*create variables of 14 types including rune, byte and also take 2 interface{} type variables;For one assign int and the other one assign float32 values.
assign values to them and perform addition and multiplication, print the details*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n1 int = 10
	var n2 int8 = 127
	var n3 int16 = 25500
	var n4 int32 = -1283000002
	var n5 int64 = 1238713891823183189
	var n6 float32 = 123341.8391923013818123913
	var n7 float64 = -126371839123.3198789123

	//var n8 bool = false
	//var n9 string = "random"

	var n10 byte = 237       //alias for int8
	var n11 rune = 838713727 //alias for int32

	var n12 float32 = 324.1234
	var n13 float32 = 255
	var n14 float32 = 7
	var n15 float32 = 87237
	var n16 float32 = 2187

	//below are the interfaces
	var add_res any = n1 + int(n2) + int(n3) + int(n4) + int(n5) + int(n6) + int(n7)
	var mul_res any = float32(n10) * float32(n11) * float32(n12) * float32(n13) * float32(n14) * float32(n15) * float32(n16)

	// fmt.Println("Addition: ", add_res, "MUL: ", mul_res)
	// var res = fmt.Sprintln(add_res, mul_res, n8, n9)

	fmt.Println(reflect.TypeOf(add_res), reflect.TypeOf(mul_res))

	//fmt.Println("Final result: ", res)

	fmt.Printf("addition result: %v, multiplication result: %v \n", add_res, mul_res)

}
