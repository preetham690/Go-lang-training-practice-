/* Create variables of all number types . 14 datatypes including rune and byte

declare and assign values of all variables

perform addition and multiplecation, assign the result to a another variable, print it. */

package main

import "fmt"

func main() {
	var n1 int = 10
	var n2 int8 = 127
	var n3 int16 = 25500
	var n4 int32 = -1283000002
	var n5 int64 = 1238713891823183189
	var n6 float32 = 123341.8391923013818123913
	var n7 float64 = -126371839123.3198789123

	var n8 bool = false
	var n9 string = "random"
	var n10 byte = 237       //alias for int8
	var n11 rune = 838713727 //alias for int32

	var n12 uint = 0
	var n13 uint8 = 255
	var n14 uint16 = 7213
	var n15 uint32 = 87237812
	var n16 uint64 = 2187318282109008380

	var add_res = n1 + int(n2) + int(n3) + int(n4) + int(n5) + int(n6) + int(n7)
	var mul_res = uint(n10) * uint(n11) * uint(n12) * uint(n13) * uint(n14) * uint(n15) * uint(n16)

	fmt.Println("Addition: ", add_res, "MUL: ", mul_res)
	var res = fmt.Sprintln(add_res, mul_res, n8, n9)

	fmt.Println("Final result: ", res)
}
