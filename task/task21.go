//create a channel
//send a value to the channel
//close the channel

//create a unconditional for loop
//use the select the statement
//needs to know how many times the closed signales

package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan int)

	go func() {
	outer:
		for {
			select {
			case v, ok := <-ch:
				fmt.Println("value from channel", v, "status of the channel", ok)
				if !ok {
					break outer
				}
			}
		}
	}()
	ch <- 100
	close(ch)

	runtime.Goexit()
}

//even when the channel is closed it will send a signal that it is closed
