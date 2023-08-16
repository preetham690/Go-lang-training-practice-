package main

import (
	"fmt"
	"os"
)

func main() {

	bytes, err := os.ReadFile("Data.txt")

	if err != nil {
		println(err)
	}
	fmt.Println("Count of chars in the file:", len(bytes))
}
