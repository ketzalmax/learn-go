package main

import "fmt"

func minimalistCaller(cb func() string) string {
	s := cb()
	return s
}

func caller(cb func(i int) string, i int) string {
	s := cb(i)
	return s
}

func calledBack(i int) string {
	if i == 0 {
		return "zero"
	} else if i == 1 {
		return "one"
	}
	return "please enter a 1 or a 0"
}

func sum(xi ...int) int {
	fmt.Printf("%T\t", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(cb func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return cb(yi...)
}

// don't actually need the identfier in the callback function
func odd(cb func(...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return cb(yi...)
}

func notDivisibleByFive(cb func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%5 != 0 {
			yi = append(yi, v)
		}
	}
	return cb(yi...)
}

func main() {
	// a super minimal anonymous callback
	supermimimal := minimalistCaller(func() string {
		return "super miminal"
	})
	fmt.Println(supermimimal)

	// an anonymous callback that meets the callback signature of function caller
	minimal := caller(func(i int) string {
		return "miminal"
	}, 1)
	fmt.Println(minimal)

	theString := caller(calledBack, 5)
	fmt.Println(theString)

	anotherCalledBack := func(i int) string {
		if i == 0 {
			return "you entered zero"
		} else if i == 1 {
			return "you entered one"
		}
		return "please enter a 1 or a 0"
	}

	theStringTwo := caller(anotherCalledBack, 1)
	fmt.Println(theStringTwo)

	ii := []int{2, 3, 4, 5, 6, 20, 30}
	s := sum(ii...)
	fmt.Println("sum of all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("callback result sum of even numbers", s2)
	s3 := odd(sum, ii...)
	fmt.Println("callback result sum of odd numbers", s3)
	s4 := notDivisibleByFive(sum, ii...)
	fmt.Println("callback result of numbers not divisible by 5", s4)

	anotherSum := func(xi ...int) int {
		fmt.Printf("%T\t", xi)
		total := 0
		for _, v := range xi {
			total += v
		}
		return total
	}

	s5 := notDivisibleByFive(anotherSum, ii...)
	fmt.Println("anotherSum callback result of numbers not divisible by 5", s5)

}
