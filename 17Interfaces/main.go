package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(nil, "Hello")
}

type Console struct{}

func (c *Console) Write(bytes []byte) (n int, err error) {
	return fmt.Println(string(bytes))
}

type FW struct {
	FileName string
}

func (f *FW) Write(bytes []byte) (n int, err error) {
	if f == nil {
		return 0, errors.New("nil  file object")
	}
	f1, err := os.Open((f.FileName))

	if err != nil {

	}
}

func write(bytes []byte) (n int, err error) {
	return 0, nil
}
