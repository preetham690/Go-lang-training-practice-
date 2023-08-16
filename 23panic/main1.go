package main

import "fmt"

func main() {
	fn := "Vs"
	ln := "&CO"

	str := GetFullName(&fn, &ln)
	fmt.Println("Full name :", *str)

	//below code generates the panic
	var fn1 string
	ln1 := "co"

	str1 := GetFullName(&fn1, &ln1)
	fmt.Println("SO the Full name is:", str1)
}

// here we are rising the panic--> in the sense we can raise our own panic
func GetFullName(fName, lName *string) *string {
	if fName == nil || *fName == "" {
		panic("first name can not be nil")
	}

	if lName == nil || *lName == "" {
		panic("last name can not be nil")
	}
	str := (*fName + *lName)
	return &str
}
