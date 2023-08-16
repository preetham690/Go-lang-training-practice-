package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for v := range GenerateSquares(100) {
			fmt.Println("rom ch1 -->", v)
		}
	}()
	go func() {
		for v := range GenerateSquares(200) {
			fmt.Println("rom ch2 -->", v)
		}
	}()

	ch3 := GenerateSquares(100)
	ch4 := GenerateSquares(200)

	go func() {
		for v := range ch3 {
			fmt.Println("from ch3 -->", v)
		}
	}()
	go func() {
		for v := range ch4 {
			fmt.Println("from ch4 -->", v)
		}
	}()
	runtime.Goexit()
}

func GenerateSquares(num int) chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < num; i++ {
			ch <- i * i
		}
		close(ch)
	}()

	return ch
}
