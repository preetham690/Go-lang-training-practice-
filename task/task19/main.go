package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ch := make(chan string)

	f, err := os.Open("data.txt")
	check(err)

	var filestr []string

	defer f.Close()

	// create a new scanner
	scanner := bufio.NewScanner(f)

	// Use scanword to split
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		filestr = append(filestr, scanner.Text())
	}

	for i, val := range filestr {
		fmt.Println("line:", i, "-", val)
		go countChars(ch)
		ch <- val
		go countWords(ch, i)
		ch <- val
		go countVowelsAndConsonants(ch, i)
		ch <- val
	}

	//var wg sync.WaitGroup

	// //number of chars
	// fmt.Println("Count of chars:")
	// wg.Add(1)
	// go func() {
	// 	countChars(ch)
	// 	wg.Done()
	// }()

	// //Identify number fo words
	// fmt.Println("Count of words:")

	// wg.Add(1)
	// go func() {
	// 	for i := 0; i < len(filestr); i++ {
	// 		countWords(ch, i)
	// 	}
	// 	wg.Done()
	// }()

	// //Identify number of vovels - Identify number of consonents - Identify number of special chars
	// fmt.Println("Count of vowels,consonants and special chars: ")
	// wg.Add(1)
	// go func() {
	// 	for i := 0; i < len(ch); i++ {
	// 		countVowelsAndConsonants(ch, i)
	// 	}
	// 	wg.Done()
	// }()

	//wg.Wait()

}

// checking the error
func check(err error) {
	if err != nil {
		panic("No data")
	}
}

// for counting the chars
func countChars(ch chan string) {
	fmt.Println("Count of characters: ", len(<-ch))
}

// this is for countwords
func countWords(ch chan string, i int) {
	st := <-ch

	split := strings.Split(st, " ")

	fmt.Printf("number of words in the sentence %v is : %v ", i, len(split))

}

// this is for checking the vowels, consonants and special chars
func countVowelsAndConsonants(ch chan string, i int) {
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
	fmt.Println("Number of Vowels in", i, "th line: ", vow, "\nNumber of consonantsin", i, "th line: ", cons)
}
