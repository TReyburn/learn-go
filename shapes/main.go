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

func main() {}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("Area: %+v", s.getArea())
}
