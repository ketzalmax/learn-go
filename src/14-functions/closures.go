package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func incrementorDecrementor() (func() int, func() int) {
	var x int
	return func() int {
			x++
			return x
		}, func() int {
			x--
			return x
		}
}

func main() {
	// creating sum here has no effect on the result, total's referenece to sum is "private"
	sum := 100
	fmt.Println(sum)
	// total has "enclosed" a reference to sum initialized at zero in adder(), and can modify it
	total := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(total(i))
	}
	// total2 gets its own new reference to sum from adder, also initialized at zero
	total2 := adder()
	for i := 0; i < 10; i++ {
		// running total2 twice in one line changes sum twice each loop
		fmt.Println(total2(i), " ", total2(i))
	}

	a := incrementor()
	fmt.Println("*** incrementor **")
	fmt.Printf("type of a %T\n", a)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("*** incrementorDecrementor ***")

	c, d := incrementorDecrementor()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(d())
	fmt.Println(d())
}
