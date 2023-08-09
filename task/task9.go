// find the biggest and smallest value in the array
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := [10]int{4, 2, 1, 5, 7, 3, 8, 0, 10, -1}

	fmt.Println(getBigAndSmall(arr))

	randArr := randomArray()

	fmt.Println(getBigAndSmall([10]int(randArr)))
}

func randomArray() []int {
	var arr [10]int

	for ind, _ := range arr {
		arr[ind] = rand.Intn(1000)
	}

	return arr[:]
}

func getBigAndSmall(arr [10]int) (big, small int) {
	big = arr[0]
	small = arr[0]

	for ind, _ := range arr {
		if arr[ind] > big {
			big = arr[ind]
		}
		if arr[ind] < small {
			small = arr[ind]
		}
	}
	return
}
