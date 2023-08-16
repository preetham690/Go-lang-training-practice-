package main

import (
	"fmt"
	"sync"
	strings "wg/string"
)

func main() {
	//wg := new(sync.Waitgroup)
	var wg sync.WaitGroup

	wg.Add(1) // if we already know how many ngoroutines are there then we can add wg directly
	go func() {

		for i := 0; i < 3; i++ {
			//n := rand.Intn(10000)
			wg.Add(1)
			go isEven(i, &wg)

		}
		wg.Add(1)
		str := rev("Preetham", &wg)
		fmt.Println(str)
		wg.Done()
	}()

	wg.Wait()
}

func isEven(n int, wg *sync.WaitGroup) {
	if n%2 == 0 {
		fmt.Println(n, "it's even")
	} else {
		fmt.Println(n, "it's odd")
	}
	wg.Done()
}

func rev(s string, wg *sync.WaitGroup) string {
	wg.Done()
	return strings.Reverse(s)

}
