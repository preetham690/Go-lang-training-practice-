package inter

import (
	"fmt"
	"i/iFace"
)

func main() {

}

func Area(ia i.IArea) { //passing an interface as a parameter. It is a form of dependency injection
	a := ia.Area()
	fmt.Println(a)
}
func Perimeter(ip i.IPerimeter) {
	p := ip.Perimeter()
	fmt.Println(p)
}
