package main

import "fmt"

// this is when the application raises a panic
// it detects the recent panic and avoids the func from raising it
// this won't work for the cascading panic

func main() {
	defer RecoverMe() // this doesn't helps you to print the str2 because the main functionj itself will crash
	fn1 := "vs"

	str1 := GetFullName(&fn1, nil)
	fmt.Println(*str1)

	ln1 := "&co"
	str2 := GetFullName(&fn1, &ln1)
	fmt.Println(*str2)
}

// when the panic is raised if look for any defered fn
func GetFullName(fName, lName *string) *string {
	//defer fmt.Println("Is there any panic") // even if everything crashes "defer will not crash"

	//defer fmt.Println("Yes there is")
	defer RecoverMe()
	if fName == nil || *fName == "" {
		panic("first name can not be nil")
	}

	if lName == nil || *lName == "" {
		panic("last name can not be nil")
	}
	str := (*fName + *lName)
	return &str
}

func RecoverMe() {
	if pan := recover(); pan != nil {
		fmt.Println(pan, "I will recover the whole stack execpt the paniced fucntion")
	}
}
