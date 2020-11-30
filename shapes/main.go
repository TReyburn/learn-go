package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

type shape interface {
	getArea() float64
}

func main() {
	s1 := square{sideLength: 10.2}
	printArea(s1)

	s2 := square{sideLength: 2}
	printArea(s2)

	t1 := triangle{
		height: 5.3,
		base:   1.1,
	}
	printArea(t1)

	t2 := triangle{
		height: 3,
		base:   2,
	}
	printArea(t2)
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("Area: %+v \n", s.getArea())
}
