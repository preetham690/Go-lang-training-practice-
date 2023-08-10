/*Find the second biggest number in an array/slice

Find the second smallest number in an array/slice

Write everything using functions.

SecondBiggest(arr []int)int

SecondSmallest(arr []int)int*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 45, 6, 6, 7, 8, 9, 9, 3, 0, 32, 3, 42, 3, 4, 23, 55}

	fmt.Println(SecondBiggest(arr))
	fmt.Println(SecondSmallest(arr))
}

func SecondBiggest(arr []int) int {
	big := arr[0]

	//this will find the last element
	for i, _ := range arr {
		if arr[i] > big {
			big = arr[i]
		}
	}

	res := arr[0] //holds the largest number

	for i, _ := range arr {
		if arr[i] > res && arr[i] != big {
			res = arr[i]
		}
	}

	return res
}

func SecondSmallest(arr []int) int {
	small := arr[0]

	//this will find the smallest element
	for i, _ := range arr {
		if arr[i] < small {
			small = arr[i]
		}
	}
	res := arr[0] //this wil hold the smallest number

	for i, _ := range arr {
		if arr[i] < res && arr[i] != small {
			res = arr[i]
		}
	}

	return res
}
