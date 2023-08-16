package main

import "fmt"

// using channels to block the code , just like waitgroups
func main() {
	ch := make(chan int)

	notify := make(chan struct{})

	go func() {
		for i := 0; i <= 100; i++ {
			ch <- i * i
		}
	}()

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("receiver -->", <-ch)
		}
		notify <- struct{}{} // Instance of struct // usally used to notify the complition
		//this notifies the channel
	}()

	<-notify // this blocks the code so that the receiver channel can respond back
}
