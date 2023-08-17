package main

import "flag"

var (
	PORT string // command line coding
)

func main() {
	flag.StringVar(&PORT, "port", "50080", "--port=50080 or -port=50080")
	flag.Parse()

}
