package main

import (
	"fmt"
	"math"
)

type square struct {
	width  int
	length int
}

type circle struct {
	radius float64
}

type shape interface {
	area()
}

func main() {
	s := square{
		width:  3,
		length: 4,
	}
	c := circle{
		radius: 5,
	}

	fmt.Println(s.area())
	fmt.Println(c.area())
}

func (s square) area() int {
	return s.length * s.width
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
