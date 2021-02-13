package main

import "fmt"

// simple defer
// func main() {
// 	defer fmt.Println("One")
// 	fmt.Println("Two")
// 	fmt.Println("Three")
// 	defer fmt.Println("Four")
// 	fmt.Println("Five")
// 	fmt.Println("Six")
// }

// *** a second defer example - within a containing function ***
// func foo() {
//	defer func() {
//		fmt.Println("foo DEFER ran")
//	}()
//	fmt.Println("foo ran")
//}

// func main() {
//	defer foo()
//	fmt.Println("Hello playground")
// }

// *** a third defer example ***
func bar(zi []int) int {
	sum := 0
	for _, v := range zi {
		sum += v
	}
	return sum
}

func deferMe(zi []int) {
	sum := 0
	for _, v := range zi {
		sum += v
	}
	fmt.Println(sum, "I was deferred first")
	// return sum, "I was deferred"
}

func main() {
	y := []int{10, 20, 30}
	z := []int{100, 200, 300}
	defer deferMe(z)
	defer fmt.Println("I was deferred second", foo(2, 3, 4))
	defer fmt.Println("I was deferred third", foo(y...))
	defer fmt.Println("I was deferred fourth", bar(z))
}
