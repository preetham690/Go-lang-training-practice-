package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var mymap map[string]any // declaration
	mymap = make(map[string]any)

	mymap["1"] = "blore-1"
	mymap["2"] = "blore-2"
	mymap["3"] = "blore-3"
	mymap["4"] = 2

	for k, val := range mymap {
		fmt.Println("Key", k, "Value", val)
	}

	val, ok := mymap["3"] //to fetch a value from the map

	if !ok {
		fmt.Println("Not found")
	} else {
		fmt.Println("It's there", val)
	}

	delete(mymap, "4") // this is how you delete the key, if present

	fmt.Println(mymap)
}

func countVal() {
	slice := []int{}

	for i, _ := range slice {
		slice[i] = rand.Intn(50)
	}

	mp := make(map[int]int)

	for i, _ := range mp {
		mp[i]
	}
}
