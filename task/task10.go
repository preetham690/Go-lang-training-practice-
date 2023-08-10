/*Take 2 dimentional array 3 X 3 and assign values .

Get the sume of Each row and each column data

Write a func to perform those operations

Input:

1 2 3

4 5 6

7 8 9
Output:

Row 1 Sum: 6

Row 2 Sum: 15

Row 3 Sum: 24

Column 1 Sum: 12

Column 2 Sum: 15

Column 3 Sum: 18*/

package main

import "fmt"

func main() {
	arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	sumOfRows(arr)
	sumOfColumns(arr)
}

func sumOfRows(arr [3][3]int) {
	for i, val := range arr {
		sum := 0
		for j, _ := range val {
			sum += arr[i][j]
		}
		fmt.Println("Row", i+1, "Sum", sum)
	}
}

func sumOfColumns(arr [3][3]int) {
	for i, val := range arr {
		sum := 0
		for j, _ := range val {
			sum += arr[j][i]
		}
		fmt.Println("Column", i+1, "Sum", sum)
	}
}
