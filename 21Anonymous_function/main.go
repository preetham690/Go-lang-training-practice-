package main

import (
	"fmt"
	"strings"
)

// Anonymous functions has same features as normal functions
// but it doesn't have a name

func main() {
	Greet()
	func() {
		str := "hello vs&co, new york"
		fmt.Println(str)
	}() //exec -> it is also called a caller

	func(str string) {
		fmt.Println(str)
	}("Hello vs&co, ohio")

	//reversing the string
	rstr := func(str string) string {
		rev := ""
		//better to use for range method to do so
		for i := 0; i < len(str); i++ { // it won't work if you use any unicode character
			rev = string(str[i]) + rev
		}
		return rev
	}("Hello vs&co, London")
	fmt.Println(rstr)

	f1 := func(str string) string { // here we are assigning the whole fucntion to the variable
		rev := ""
		for _, v := range str {
			rev = string(v) + rev
		}
		return rev
	} //here we are not using the executor

	out := f1("Hi, myself Preetham")
	fmt.Println(out)

	//anonymous fn to return count of vovels , consonents and special chars
	f2 := func(str string) (v, c, s int) {
		vov := "aeiou"
		cons := " bcdfghjklmnpqrstvwxyz"
		for i := range str {
			//for vovels
			if strings.Contains(vov, string(str[i])) {
				v++
			} else if strings.Contains(cons, string(str[i])) {
				c++
			} else {
				s++
			}
		}
		return
	}
	count1, count2, count3 := f2("hi, myself preetham")
	fmt.Println("Vovels:", count1, "consonants:", count2, "Special chars:", count3)

	//function closure
	sum := Calc(10, 20, func(a1, b1 int) int { //inline
		return a1 + b1
	})
	fmt.Println("sum is:", sum)

	//here it is done by assigning fn to a variable
	fnm := func(a, b int) int {
		return a * b
	}
	mul := Calc(10, 20, fnm)
	fmt.Println(mul)

	Operations()
}

func Greet() {
	fmt.Println("Hi, I'm from VS&Co, Blore")
}

// closure
func Calc(a, b int, fn func(a1, b1 int) int) int { //here we are passing function as a argument
	return fn(a, b)
}

func Operations() {
	m := make(map[string]func(int, int) int)
	m["add"] = func(i1, i2 int) int { return i1 + i2 }
	m["sub"] = func(i1, i2 int) int { return i1 - i2 }
	m["mul"] = func(i1, i2 int) int { return i1 * i2 }
	m["div"] = func(i1, i2 int) int { return i1 % i2 }

	//fmt.Println("sum :", m["add"](10, 20))

	//here we are running the same value for the all the keys
	for i, val := range m { // here the value is function
		r := val(100, 200)
		fmt.Println("key:", i, "fn as value:", r)
	}
}
