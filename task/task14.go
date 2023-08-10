package main

func main() {

}

type Company struct {
	id, telephone, fax int
	name, website      string
	address            // promoted
}

type address struct {
}
