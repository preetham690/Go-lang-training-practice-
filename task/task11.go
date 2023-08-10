/*Create and assign values to 2d array 3X3 array
store rows to columns and columns to rows in the array
Write a func to perform those operations
Input:

1 2 3

4 5 6

7 8 9
Output:

 1 4 7

 2 5 8

 3 6 9*/

package main

import "fmt"

func main() {
	arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	ans := transpose(arr)
	printState(ans)

}

func transpose(arr [3][3]int) [3][3]int {
	var ans [3][3]int
	for i, val := range arr {
		for j, _ := range val {
			ans[j][i] = arr[i][j]
		}
	}
	return ans
}

func printState(arr [3][3]int) {
	for i, val := range arr {
		for j, _ := range val {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
