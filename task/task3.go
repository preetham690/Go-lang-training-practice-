//Create constants and perform mathematical calc on constants

package main

import (
	f "fmt"
	m "math"
	"reflect"
)

const pi = m.Pi

var temp float32 = pi + pi

func main() {

	f.Printf("PI value is: %v and the type is %v\n", pi, reflect.TypeOf(pi))

	//addition
	f.Println("Add: ", pi+pi)

	//substraction
	f.Println("Sub: ", pi-pi)

	//mul
	f.Println("Mul: ", pi*pi)

	//division
	f.Println("Div: ", pi/pi)

	//modulous
	f.Println("Mod: ", m.Mod(pi, pi))
}
