package main

import "fmt"

func main() {
	var iCal ICalculator

	//this is for int type
	ic1 := &IntType{A: 20, B: 30}
	iCal = ic1

	r1 := iCal.add()
	fmt.Println("Sum for the type int: ", r1)

	//this is for float32
	ic2 := &Float32Type{A: 10, B: 10}
	iCal = ic2

	r2 := iCal.add()
	fmt.Println("for float32: ", r2)

	//this is for float64
	ic3 := &Float64Type{A: 20.1, B: 30.9}
	iCal = ic3

	r3 := iCal.add()
	fmt.Println("for float64: ", r3)

	//this is for string
	ic4 := &String{A: "10", B: "20"}
	iCal = ic4

	r4 := iCal.add()
	fmt.Println("sum of two string numbers:", r4)

	//--> this is for substraction

	ic5 := &IntType{A: 20, B: 20}
	iCal = ic5

	r5 := iCal.sub()
	fmt.Println("sub res for type int: ", r5)
}

type ICalculator interface {
	add() any
	sub() any
	// mul()
	// div()
	// get()
	// clear()
}

type IntType struct {
	A, B int
	res  int
}
type Float32Type struct {
	A, B float32
	res  float32
}
type Float64Type struct {
	A, B float64
	res  float64
}

// sub implements ICalculator.
func (*Float64Type) sub() any {
	panic("unimplemented")
}

type String struct {
	A, B string
	res  string
}

func (i *IntType) add() any {
	i.res = i.A + i.B
	return i.res
}
func (f *Float32Type) add() any {
	f.res = f.A + f.B
	return f.res
}
func (f *Float64Type) add() any {
	f.res = f.A + f.B
	return f.res
}
func (s *String) add() any {
	s.res = s.A + s.B
	return s.res
}

// type () dispaly() {
// 	fmt.Println("")
// }

// for substraction
func (i *IntType) sub() any {
	i.res = i.A - i.B
	return i.res
}
func (f *Float32Type) sub() any {
	f.res = f.A - f.B
	return f.res
}
func (s *String) sub() any {
	return s.res
}
