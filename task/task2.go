//Use fmt package functions like Print, Println, Printf and identify the differences

package main

import "fmt"

func main() {
	var name string = "Preetham"
	var age int = 22

	//printf -> Print formatter

	fmt.Printf("Hi this is, %v and my age is %v\n", name, age)

	//print -> takes the string and print it

	fmt.Print("Hi this is", name, "and my age is", age, "\n")

	//println -> add a new line at the end along with the spaces inbetween the variables
	fmt.Println("Hi this is", name, "and my age is", age)

	//Sprint -> the thing is "go" does not supports concatination of multiple types

	var sen string = fmt.Sprintf("Hi, Myself %v and I'm %v years old", name, age)

	fmt.Println(sen)
}
