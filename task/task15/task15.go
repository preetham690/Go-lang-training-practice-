//Implement Circle for example 33-interfaces

package main

import (
	"fmt"
	"taskOnInter/interfaces"
	"taskOnInter/shape"
)

func main() {
	r1 := &shape.Circle{R: 20}

	var s1 shape.Circle = 24.9

	slice3 := make([]interfaces.IAreaPerimeter, 0)

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
