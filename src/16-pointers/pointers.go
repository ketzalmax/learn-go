package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println("memory address of a: ", &a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a // *int type is a pointer to an int
	fmt.Println(b)

	*b = 43 // b holds the address of a, dererence it and set that value to 43
	fmt.Println(a)

	c := &a // type inference ??
	fmt.Println(c)
	fmt.Println(*c)  // * OPERATOR shows the value stored at that pointer/address, aka "dereferences"
	fmt.Println(*&a) // get address at &a and dereference it
	// var d int = &a // this would not work, int and *int a different types

}
