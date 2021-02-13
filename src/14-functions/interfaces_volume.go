package main

import (
	"fmt"
	"math"
)

type ofStructure interface {
	volume() float64
}

type triangularPyramid struct {
	bs1    float64
	bs2    float64
	bs3    float64
	height float64
}

type cylinder struct {
	height float64
	radius float64
}

type box struct {
	length float64
	width  float64
	height float64
}

type sphere struct {
	radius float64
}

type cube struct {
	length float64
}

func (tp triangularPyramid) volume() float64 {
	s := (tp.bs1 + tp.bs2 + tp.bs3) / 2
	ba := math.Sqrt((s - tp.bs1) * (s - tp.bs2) * (s - tp.bs3))
	// return 1 / 3 * (ba * tp.height)
	return (ba * tp.height) / 3
}

func (c cylinder) volume() float64 {
	return (math.Pi * (c.radius * c.radius) * c.height)
}

func (c cube) volume() float64 {
	return c.length * c.length * c.length
}

func (s sphere) volume() float64 {
	return (4 * math.Pi * math.Pow(s.radius, float64(3))) / 3
}

func (b box) volume() float64 {
	return b.length * b.width * b.height
}

func calculateVolume(kind ofStructure, shapeType string) {
	fmt.Printf("The volume of the %s is: %f \n", shapeType, kind.volume())
}

func main() {
	tp := triangularPyramid{
		bs1:    10,
		bs2:    15,
		bs3:    20,
		height: 100,
	}

	c := cube{
		length: 7,
	}

	b := box{
		length: 5.5,
		width:  5.5,
		height: 7.7,
	}

	s := sphere{
		radius: 7.14,
	}

	cyl := cylinder{
		height: 10,
		radius: 7,
	}

	calculateVolume(c, "Cube")
	calculateVolume(b, "Box")
	calculateVolume(s, "Sphere")
	calculateVolume(cyl, "Cylinder")
	calculateVolume(tp, "Triangular Pyramid")

}
