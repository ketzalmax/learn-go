package main

import "fmt"

func returnFunc() func() string {
	return func() string {
		return "I was returned"
	}
}

func main() {

	x := returnFunc()
	fmt.Println(x())
	fmt.Printf("type of x: %T\n", x)
	fmt.Println("x executed", x())
	fmt.Println("returnFunc, immediately executed", returnFunc()())

	y := x()
	fmt.Println("y contains return of x executed: ", y)

	// anonymous func example
	f := func() func() string {
		return func() string {
			return "My name is Bruce"
		}
	}()

	s := f()
	fmt.Println(s)
}
