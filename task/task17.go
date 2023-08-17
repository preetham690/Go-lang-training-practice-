package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(ch chan string) {
		str := <-ch
		var vow, cons, specialchr = 0, 0, 0
		for _, val := range str {
			if val == 65 || val == 69 || val == 73 || val == 79 || val == 85 || val == 97 || val == 101 || val == 105 || val == 111 || val == 117 {
				vow++
			} else if val >= 65 && val <= 91 || val >= 97 && val <= 121 {
				cons++
			} else {
				specialchr++
			}
		}

		fmt.Println("Number of Vowels in", vow)
		fmt.Println("Number of consonantsin", cons)
		fmt.Println("Number of Special chars", specialchr)
		wg.Done()
	}(ch)

	ch <- "preetham shetty"
	close(ch)

	wg.Wait()
}
