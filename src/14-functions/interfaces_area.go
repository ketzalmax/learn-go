package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	name() string
}

type triangle struct {
	sideA float64
	sideB float64
	sideC float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (t triangle) area() float64 {
	s := (t.sideA + t.sideB + t.sideC) / 2
	return math.Sqrt((s - t.sideA) * (s - t.sideB) * (s - t.sideC))
}

func (c circle) area() float64 {
	return (c.radius * c.radius) * math.Pi
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (t triangle) name() string {
	return "Triangle"
}

func (r rect) name() string {
	return "Rectangle"
}

func (c circle) name() string {
	return "Circle"
}

func getArea(s shape) {
	fmt.Println("The area of the unknown shape is:", s.area())
}

func getAreaAndName(s shape) {
	fmt.Printf("The area of the %v is: %v\n", s.name(), s.area())
}

func getArea2(s shape) {
	fmt.Printf("The type of the shape is %v and the area is %v\n", s.name(), s.area())
}

func main() {
	rect1 := rect{
		width:  5,
		height: 4,
	}

	circle1 := circle{
		radius: 5,
	}

	triangle1 := triangle{
		sideA: 5,
		sideB: 6,
		sideC: 7,
	}

	// using interface
	getArea(rect1)
	getArea(circle1)
	getAreaAndName(triangle1)
	getAreaAndName(rect1)
	getArea2(rect1)

	// using methods
	fmt.Println("The area of the rect is", rect1.area())
	fmt.Println("The area of the circle is", circle1.area())
}
