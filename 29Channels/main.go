package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int) // unbuffered channel
	wg := new(sync.WaitGroup)
	// ch <- 100 // assigning a value to a channel
	// // here the main is blocked because main sends the val and recieves a value
	// num := <-ch // recieve a value from channel

	wg.Add(1)
	go func() {
		num := <-ch
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				fmt.Println(num)
			}
			wg.Done()
		}()

		wg.Done()
	}()

	time.Sleep(time.Millisecond * 1000)
	ch <- 100
	wg.Wait()
	//fmt.Println(num)
}
