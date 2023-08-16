package main

import "fmt"

func main() {
	ch := make(chan int)
	notify := make(chan struct{})

	go receiver(ch, notify)
	go sender(ch)
	<-notify //this is notifying that to close the channel//channel has been closed
}

func sender(ch chan int) {
	for i := 0; i < 101; i++ {
		ch <- i * i
	}
}

func receiver(ch <-chan int, notify chan<- struct{}) { // "<-" meaning sender
	for i := 0; i < 101; i++ {
		fmt.Println("received -->", <-ch)
	}
	notify <- struct{}{} //this sends the notification
}

//we can close a channel using "close(ch)"
//only sender can close a channel
//we can use range loop to traverse through the channels, but the thing is we need to close the channel
