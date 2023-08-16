//Implement Circle for example 33-interfaces

package main

import (
	"fmt"
	interfaces "taskOnInter/Interface"
	"taskOnInter/shape"
)

func main() {
	r1 := &shape.Circle{R: 20}
	fmt.Println(r1)

	var s1 shape.Circle
	s1.R = 12.43

	slice1 := make([]interfaces.IAreaPerimeter, 0)
	slice1 = append(slice1, r1, &s1)

	for _, iface := range slice1 {
		AreaAndPetimeter(iface)
	}
}

func Area(iarea interfaces.IArea) { // interface as a parameter.It can be a form of dependency injecttion
	a := iarea.Area()
	fmt.Println("Area :", a)
}

func Perimeter(iperimeter interfaces.IPerimeter) {
	p := iperimeter.Perimeter()
	fmt.Println("Perimeter:", p)
}
func AreaAndPetimeter(iAreaPerimeter interfaces.IAreaPerimeter) {
	iAreaPerimeter.Display()
	Area(iAreaPerimeter)
	Perimeter(iAreaPerimeter)
	fmt.Println("----------------")
}
