package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}

	arr1 := [3]int{12, 13, 14}

	slice2 := arr1[:] // array to slice

	slice3 := arr1[1:] // coping a part of the array using the ":"

	//appending 2 slices together

	slice1 = append(slice1, slice2...)

	fmt.Println(slice1, slice2, slice3)

	//generally we use make -> keyword to create a slice
	//make is used for *maps, *slices and channels

	//slice4 := make([]int, 3)

	slice5 := make([]int, 6, 10) //(type , current size , capacity)
	//here in the slice we can increase the capacity by appending the value/slice to that particular slice

	slice5[0], slice5[1], slice5[2], slice5[3], slice5[4], slice5[5] = 1, 2, 3, 4, 5, 6

	fmt.Println(slice5)

	//slices are reference types , so if you make any changes in either of the slice . it will be reflected

	a := []int{1, 2, 3}
	b := []int{5, 6, 7}

	//shallow copy

	temp1 := a

	//deep copy -> is also called clone

	temp2 := make([]int, len(b))
	copy(temp2, b)

	fmt.Println(temp1, temp2)

}
