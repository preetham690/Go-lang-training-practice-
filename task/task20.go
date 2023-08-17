package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := VariadicChan(c1, c2, c3)
	c5 := VariadicChan(c1, c2)

	go func() {
		c1 <- 10
		c2 <- 20
		c3 <- 30
	}()

	go func() {
		c1 <- 10
		c2 <- 20

	}()

	// for val := range c1 {
	// 	fmt.Println(val)
	// }

	for val := range c4 {
		fmt.Println("Sum of values from all channels:", val)
	}
	for val := range c5 {
		fmt.Println("Sum of Values from all channels: ", val)
	}

}

func VariadicChan(chs ...chan int) chan int {
	ch := make(chan int)

	go func() {
		sum := 0
		for _, val := range chs {
			sum += <-val
		}
		//ch <- <-chs[0] + <-chs[1] + <-chs[2]
		ch <- sum
		close(ch)

	}()

	return ch
}
