package main

import (
	"example/example/shapes"
	"fmt"
	//_ "https://github.com/preetham690/example/tree/86e24db1ff7457e6d4f8f1ca7073e53157dd4df0/shapes"
)

func main() {
	r1 := shapes.React{L: 10.34, B: 13.45}
	a1 := r1.Area()

	fmt.Println("Area:", a1)
}
