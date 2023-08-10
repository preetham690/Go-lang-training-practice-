package main

import "fmt"

func main() {
	//type inference of arrays
	var arr [5]int
	var arr1 [5]int

	//for-range
	for i, _ := range arr {
		arr[i] = i
	}
	fmt.Println(arr)

	arr = arr1

	sum := 0

	for i, _ := range arr1 {

		arr1[i] = sum
		sum += i
	}

	fmt.Println(arr1)

	multiDArray()
}

// Multi dementational array
func multiDArray() {
	arr := [3][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9, 3, 7}}

	sum := 0
	for i, val := range arr {
		for j, _ := range val {
			sum += arr[i][j]
		}
	}

	fmt.Println("AND the sum is: ", sum)
}
