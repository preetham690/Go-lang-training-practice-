package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	//ch <- 30 //--> this will block the channel because the capacity is only 2
	fmt.Println(<-ch, <-ch) //this is a receiver

	ch <- 30
	fmt.Println(<-ch)
}
