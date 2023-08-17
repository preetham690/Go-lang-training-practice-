package main

import (
	"fmt"
	"strings"
)

func main() {
	// ch1 := make(chan string)
	// ch2 := make(chan string)
	// ch3 := MultipleChanOut(ch1, ch2)

	// go func() {
	// 	ch1 <- "vs"
	// 	ch2 <- "& co"
	// }()

	// for val := range ch3 {
	// 	fmt.Println(val)
	// }

	// ch4 := GenerateSquares(10)
	// for val := range ch4 {
	// 	fmt.Println(val)
	// }

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := VariadicChan(c1, c2, c3)

	go func() {
		c1 <- 10
		c2 <- 20
		c3 <- 30
	}()

	for val := range c4 {
		fmt.Println(val)
	}
	//fmt.Println(c4)
}

// this is fan-in pattern in channels
func MultipleChanOut(ch1, ch2 chan string) chan string {
	ch := make(chan string)

	go func() {
		ch <- <-ch1 + " " + <-ch2

		close(ch)
	}()

	return ch
}

// this is fan-out pattern in channels
func SingleToMultiple(ch1 chan string) (ch2, ch3 chan string) {
	ch2, ch3 = make(chan string), make(chan string)
	go func() {
		val := <-ch1

		vals := strings.Split(val, " ") // assuming that there's only space

		ch2 <- vals[0]
		ch3 <- vals[1]
		close(ch2)
		close(ch3)
	}()
	return
}
func GenerateSquares(num uint) chan int {
	ch := make(chan int)

	go func() { // This is goroutine, so it will be executed as a separate fn in bg
		for i := 0; i < int(num); i++ {
			ch <- i * i
		}
		close(ch)
	}()

	return ch
}

func VariadicChan(chs ...chan int) chan int {
	ch := make(chan int)

	go func() {
		ch <- <-chs[0] + <-chs[1] + <-chs[2]
		// for i := 0; i < len(chs); i++ {
		// 	close(chs[i])
		// }
		close(ch)

	}()

	return ch
}
