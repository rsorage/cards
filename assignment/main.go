package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	width float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.width * s.width
}

func printArea(s shape) {
	fmt.Println("The area of the shape is:", s.getArea())
}

func main() {
	s := square{width: 2.25}
	t := triangle{base: 1.34, height: 8.73}

	printArea(s)
	printArea(t)
}
