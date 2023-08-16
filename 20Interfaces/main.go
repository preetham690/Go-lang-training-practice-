package main

import "fmt"

// chain of actions
func main() {
	//var iCal ICalculator
	iCal := new(Notype)

	r1 := iCal.add(10).add(20).add(30).add(40).get()
	fmt.Println("The sum for r1 is: ", r1)

	r2 := iCal.add(200).sub(12).mul(571).get()
	fmt.Println("the result of r2 is:", r2)
}

type ICalculator interface {
	add(any) ICalculator
	sub(any) ICalculator
	mul(any) ICalculator
	get() float64
}

type Notype struct {
	res float64
}

func (i *Notype) add(n any) ICalculator {
	switch n.(type) {
	case int:
		i.res = i.res + float64(n.(int))
	case int8:
		i.res = i.res + float64(n.(int8))
	case int16:
		i.res = i.res + float64(n.(int16))
	case int32:
		i.res = i.res + float64(n.(int32))
	case int64:
		i.res = i.res + float64(n.(int64))
	case float32:
		i.res = i.res + float64(n.(float32))
	}
	return i
}

func (i *Notype) sub(n any) ICalculator {
	switch n.(type) {
	case int:
		i.res = i.res - float64(n.(int))
	case int8:
		i.res = i.res - float64(n.(int8))
	case int16:
		i.res = i.res - float64(n.(int16))
	case int32:
		i.res = i.res - float64(n.(int32))
	case int64:
		i.res = i.res - float64(n.(int64))
	case float32:
		i.res = i.res - float64(n.(float32))
	}
	return i
}

func (i *Notype) mul(n any) ICalculator {
	switch n.(type) {
	case int:
		i.res = i.res * float64(n.(int))
	case int8:
		i.res = i.res * float64(n.(int8))
	case int16:
		i.res = i.res * float64(n.(int16))
	case int32:
		i.res = i.res * float64(n.(int32))
	case int64:
		i.res = i.res * float64(n.(int64))
	case float32:
		i.res = i.res * float64(n.(float32))
	}
	return i
}

func (i *Notype) get() float64 {
	return i.res
}
