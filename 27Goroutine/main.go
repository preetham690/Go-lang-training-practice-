package main

import (
	"fmt"
	"runtime"
)

func main() {
	//go fmt.Println("Hello vs & co")
	// count := 0
	// go func() {
	// 	for i := 0; i < 1000; i++ {
	// 		n := rand.Intn(10000)
	// 		go isEven(n)
	// 		count++
	// 	}
	// }()
	//go runtime.Goexit() //what is happening here is the main fn is not waiting for any other goroutines
	//the only solution is to use waitgroups

	//defer fmt.Println("Hi vsandcojhvagjdajdjhasbjdbaksdbajdbakj")
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println("hi")
			go isEvenOrOdd(i)
			runtime.Goexit()

		}

	}()
	//runtime.Goexit()
}

func isEven(n int) {
	if n%2 == 0 {
		fmt.Println(n, "it's even")
	} else {
		fmt.Println(n, "it's odd")
	}
}

func isEvenOrOdd(n int) {
	if n%2 == 0 {
		fmt.Println(n, "it's even")
	} else {
		fmt.Println(n, "it's odd")
	}
}
