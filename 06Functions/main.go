package main

import "fmt"

func main() {
	greet()
	greetings("preeth")
	fmt.Println(area(2.3, 3.4))

	fmt.Println(isosceles_Triangle(2, 3))

	num := 100
	increment(&num)

	fmt.Println(num)
}

func greet() {
	// a := 1
	// count := 10
	// for count != a {
	// 	fmt.Println("Hello, bye")
	// 	count--
	// }

	fmt.Println("HI")
}

func greetings(name string) {
	fmt.Println("Hi", name)
}

func area(l, b float32) float32 {
	return l * b
}

func isosceles_Triangle(b, h float32) (area, perimeter float32) {
	return .5 * b * h, 2*b + h
}

// pass by ref
func increment(num *int) {
	*num++
}
