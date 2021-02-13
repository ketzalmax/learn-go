package main

import "fmt"

func main() {
	n, _ := fmt.Println("hello world", 42, true)
	fmt.Println("number of bytes written?", n)
	// fmt.Println("was there an error?", x)

	afunc()
}

func afunc() {
	fmt.Println("I'm in afunc")
}
