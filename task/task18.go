package main

import "fmt"

func main() {
	fmt.Println(Calc("Hi", "Victoria"))
	fmt.Println(Calc([]int{10, 12}, 10))
	fmt.Println(Calc(10, 20))
}

func Calc(a, b any) float64 {
	defer RecoverMe()
	sum := 0.0
	switch a.(type) {
	case int:
		sum = float64(a.(int) + b.(int))
	case int8:
		sum = float64(a.(int8) + b.(int8))
	case int16:
		sum = float64(a.(int16) + b.(int16))
	case int32:
		sum = float64(a.(int32) + b.(int32))
	case int64:
		sum = float64(a.(int64) + b.(int64))
	case uint:
		sum = float64(a.(uint) + b.(uint))
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
	case uint16:
		sum = float64(a.(uint16) + b.(uint16))
	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
	default:
		panic("can not perform addition")
	}
	return sum
}

func RecoverMe() {
	if pan := recover(); pan != nil {
		fmt.Printf("%v \npanic is recovered \n", pan)
	}
}
