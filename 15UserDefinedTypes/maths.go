package main

import "fmt"

func main() {
	r1 := Rect{L: 20, B: 30}

	fmt.Println(r1.area())
	fmt.Println(r1.perimeter())

	r2 := new(Rect)

	fmt.Println(r2.area())
}

type Rect struct {
	L, B float32
	A, P float32
}

type Square struct {
	S float32
}

// to write a method, need to add a receiver
// recivers are call by value, unless they are references
// these are not only for stucts
func (r *Rect) area() float32 {
	r.A = r.L * r.B
	return r.A
}
func (r *Rect) perimeter() float32 {
	r.A = 2 * (r.L * r.B)
	return r.A
}

// this is a function
func Area_Of_Rect(r Rect) float32 {
	return r.L * r.B
}

func Perimeter_Of_Rect(r Rect) float32 {
	return 2 * (r.L * r.B)
}
