package main

import "fmt"

// this is similar to boxing
func main() {
	var iface any //it will hold nil in the beginning

	iface = 1000

	var f1 float32 = 10.22

	iface = f1

	var ok1 bool = true
	iface = ok1
	str1 := "VS and co"
	iface = str1

	str2 := iface.(string) //this is called type assersion

	fmt.Println(iface)
	fmt.Println(str2)
}
