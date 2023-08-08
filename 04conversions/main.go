package main

import "fmt"

func main() {

	//HERE WE ARE DOING TYPE CONVERTIONS / CASTING
	//There is no implicit casting in GO

	var a int = 65001
	var b uint8 = uint8(a)

	fmt.Println(b) //here it prints the 8bits of the variable "a"

	var str1 string = "H"
	var n1 int = int(str1[0])

	fmt.Println(n1)

	var n2 int = 65
	var str2 string = fmt.Sprint(n2) //we can use sprint to print as it is

	fmt.Println(str2)

	ok := false

	str3 := fmt.Sprintln(ok, str1)

	fmt.Println(str3)
}
