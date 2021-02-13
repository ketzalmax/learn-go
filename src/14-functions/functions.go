package main

import "fmt"

// func (r receiver) identifier(parameters) (return(s)) {code}

// ** functions lesson 1 ***
// func main() {
// 	foo()

// 	bar("the argument.")

// 	s1 := moo()
// 	fmt.Println(s1)

// 	// s1 := "this is a string"
// 	s2 := woo("this will be returned")
// 	fmt.Println(s2)

// 	x, y, z := multiplereturn("James", 32, true)
// 	fmt.Println(x, "is ", y, " years old, and he is ", z)
// }

// func foo() {
// 	fmt.Println("an argumentS")
// }

// func bar(thestringparameter string) {
// 	fmt.Println("the parameter is now: ", thestringparameter)
// }

// func moo() string {
// 	// return fmt.Sprint("Hello from moo,", 1+1)
// 	return "Hello from moo"
// }

// func woo(anotherstring string) string {
// 	return fmt.Sprint("Hello from woo,", anotherstring)
// }

// func multiplereturn(name string, age int, ismale bool) (string, int, bool) {
// 	a := fmt.Sprintln("name is:", name)
// 	b := age
// 	c := ismale
// 	return a, b, c

// }

// func main() {
// 	xi := []int{2, 3, 4, 5, 6, 7, 8}

// 	sumSpread := sum(xi...)
// 	// sum := sum(20, 30, 40, 50, 60)
// 	// fmt.Println("The total is", sum)
// 	fmt.Println("The sumspread total is", sumSpread)
// 	ss := []string{"rat", "cat", "bat", "mat"}

// 	f(12, "whatever", ss...)
// }

// func sum(x ...int) int {
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)

// 	sum := 0
// 	for i, v := range x {
// 		sum += v
// 		fmt.Println("at index position ", i, "adding", v, "and sum is", sum)
// 	}
// 	fmt.Println("The total is", sum)
// 	return sum
// }

// func f(m int, s string, p ...string) {
// 	// do something with the stuff passed in
// }

// *** methods and interfaces ***
// Alternate way to attach methods https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922234#questions/6780676

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first, sa.last, "speak from secretAgent")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "speak from person")
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Jar",
			"Head",
		},
		ltk: true,
	}

// 	p1 := person{
// 		first: "Dr.",
// 		last:  "Yes",
// 	}

// 	fmt.Println(sa1)
// 	sa1.speak()
// 	sa2.speak()
// 	fmt.Println(p1)

// 	bar(sa1)
// 	bar(sa2)
// 	bar(p1)
// }

// *** Anonymous function and function expressions ***
// var w = 100

// func main() {
// 	// anonymous func --only inside another fucn
// 	// y could be coming in from the outside
// 	z := 99
// 	func(x int, y int, s string) {
// 		fmt.Println("THe numbers are", x, y, s)
// 	}(w, z, "yes they are")

// 	// function expression
// 	f := func() {
// 		fmt.Println("a function expression, f is now a pointer to the function")
// 	}
// ()	f()

// 	g := func(x int) {
// 		fmt.Println("the year big brother started watching", x)
// 	}
// 	g(1984)
// }

// func notDivisibleByFive(cb func(xi ...int) int, vi ...int) int {
// 	var yi []int
// 	for _, v := range vi {
// 		if v%5 != 0 {
// 			yi = append(yi, v)
// 		}
// 	}
// 	return cb(yi...)
// }



// *** recursive function ***
// func main() {
// 	// n := factorial(4)
// 	// fmt.Println(n)
// 	nf := loopFactorial(5)
// 	fmt.Println(nf)
// }

// func factorial(n int) int {
// 	// fmt.Println("n is", n)
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

// func loopFactorial(n int) int {
// 	fmt.Println("n is", n)
// 	total := 1
// 	for ; n > 0; n-- {
// 		total *= n
// 	}
// 	return total
// }