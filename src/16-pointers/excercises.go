package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	fmt.Println(p.last)
	fmt.Println((*p).last)
	(*p).last = "Changed"
}

func main() {
	person1 := person{
		first: "James",
		last:  "Bond",
	}

	changeMe(&person1)
	fmt.Println(person1.last)
}
