/*Create Company struct Id,Name, Website, Telehone, Fax, Address // Address must be a promoted filed

Create Address struct City, Line1, Street, State, Country, PinCode , Map // Map must be a promoted field

Create Map struct Lan, Lat

-> Can you access Lan and Lap with the object of Company*/

package main

import "fmt"

func main() {
	c1 := Company{1, 1234567, 830000, "VS&CO", "victoriasecret.com", address{"blore", "q2", "bsk", "Karnataka", "India", 123456, Map{1234567, 98765}}}
	fmt.Println(c1.address.Map.Lan)

	c2 := Company{2, 6734567, 840000, "PINK", "PINK.com", address{"columbus", "a1", "abc", "ohio", "usa", 98219, Map{Lan: 45678, Lat: 256789}}}

	fmt.Println(c2.address.Map)
}

type Company struct {
	id, telephone, fax int
	name, website      string
	address            // promoted
}

// This is promoted
type address struct {
	city, Line1, Street, State, Country string
	PinCode                             int
	Map                                 // promoted
}

// This is promoted
type Map struct {
	Lan, Lat int
}
