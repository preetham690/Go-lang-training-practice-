package main

import (
	"errors"
	"fmt"
)

func main() {
	var mymap = make(map[int]string)

	mymap[1] = "blore"
	mymap[2] = "mysore"
	mymap[3] = "kodagu"
	mymap[4] = "hassan"

	fmt.Println(isDeleted(mymap, 5))
	fmt.Println(update(mymap, 4, "mlore"))
}

func isDeleted(mp map[int]string, key int) (bool, error) {
	if mp == nil {
		return false, errors.New("Map Error")
	}
	_, ok := mp[key]

	if ok {
		delete(mp, key)
		return ok, nil
	} else {
		return ok, nil
	}
}

func update(mp map[int]string, key int, val string) (bool, error) {
	if mp == nil {
		return false, errors.New("Nil Map")
	}

	_, ok := mp[key]

	if ok {
		mp[key] = val
		return ok, nil
	} else {
		return ok, errors.New("key does not exist hence cannot Update")
	}
}
