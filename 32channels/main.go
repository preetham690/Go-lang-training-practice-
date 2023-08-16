package main

import "fmt"

func main() {
	ch := make(chan int)

	notify := make(chan struct{})
	notifyclosed := make(chan bool)

	go func() {
		notifyclosed <- false
	}()
	go receiver(ch, notify, notifyclosed)
	go sender(ch, notifyclosed)
	go sender(ch, notifyclosed)
	<-notify
}

func sender(ch chan int, notifyclosed <-chan bool) {
	for i := 0; i < 10; i++ {
		ok := <-notifyclosed
		if !ok {
			ch <- i * i
		} else {
			fmt.Println("it's closed")

		}
	}
	close(ch)
}

func receiver(ch <-chan int, notify chan struct{}, notifyclosed chan<- bool) {
	for {
		val, ok := <-ch

		if !ok {
			notifyclosed <- true
			break
		} else {
			notifyclosed <- false
			fmt.Println("Received <-- ", val)
		}
	}
	notify <- struct{}{}
}
