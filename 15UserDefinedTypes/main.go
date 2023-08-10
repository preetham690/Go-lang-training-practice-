package main

import "fmt"

func main() {
	p1 := Person{id: 1, name: "Preet", email: "pre@abc.com", mob: 1234568, addr: address{street: "bsk", city: "blore", state: "karnataka", country: "India", pincode: 12345}}

	e1 := employee{id: 1, name: "Preet", email: "pre@abc.com", mob: 1234568}
	e1.city = "goa"
	e1.address.country = "India"

	var p2 Person

	p2.id = 2
	p2.name = "Paarth"
	p2.email = "Parth@vs.com"
	p2.mob = 8127432

	fmt.Println(p1, p2)

	//this is anonymous struct
	as1 := struct {
		name, email string
	}{
		name:  "abc",
		email: "abc.@com",
	}

}

// this comes under user defined data
type Person struct {
	id    int
	name  string
	email string
	mob   int
	addr  address // composition -> embedding another struct here
}
type address struct {
	street, city, state, country string
	pincode                      int
}
type employee struct {
	id      int
	name    string
	email   string
	mob     int
	address // promoted fields -> they don't have the names
}

// aslo we can create inline embedded struct
