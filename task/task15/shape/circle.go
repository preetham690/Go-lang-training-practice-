package shape

import "fmt"

type Circle struct {
	R float32
}

func (c *Circle) Area() float32 {
	return 3.14 * c.R * c.R
}

func (c *Circle) Perimeter() float32 {
	return 2 * 3.14 * c.R
}
func (c *Circle) Display() {
	fmt.Println("Circle")
	fmt.Println("-----------------")
}
