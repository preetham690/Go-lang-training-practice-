package main

import (
	"example/example/shapes"
	"fmt"
)

func main() {
	r1 := shapes.React{L: 10.34, B: 13.45}
	a1 := r1.Area()

	fmt.Println("Area:", a1)
}
