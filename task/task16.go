package main

import "fmt"

func main() {
	//this is for slice
	slice := StringSlice{"a", "b", "c", "d"}

	var sliceInter interExample

	sliceInter = slice
	fmt.Println("To-string:", sliceInter.ToString())
	fmt.Println("To-count:", sliceInter.ToCount())

	//this is for maps

	m := StringMap{"1": "task1", "2": "task2", "3": "task3"}

	var mapInter interExample

	mapInter = m
	fmt.Println("To-String:", mapInter.ToString())
	fmt.Println("To-Count:", mapInter.ToCount())

}

type interExample interface {
	ToString() string
	ToCount() int
}

// for slice
type StringSlice []string

func (ss StringSlice) ToString() string {
	str := ""
	for i := range ss {
		str += ss[i] + " "
	}
	return fmt.Sprintf("%v", str)
}

func (ss StringSlice) ToCount() int {
	return len(ss)
}

//for maps

type StringMap map[string]string

func (sm StringMap) ToString() string {
	str := ""

	for i, val := range sm {
		str += i + ":" + val + " "
	}
	return fmt.Sprintf("%v", str)
}

func (sm StringMap) ToCount() int {
	return len(sm)
}
