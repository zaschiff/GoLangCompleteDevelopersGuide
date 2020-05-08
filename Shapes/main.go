package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	testS := square{sideLength: 2.0}
	testT := triangle{height: 2.0, base: 3.0}

	printArea(testS)
	printArea(testT)

}

func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}

func (t triangle) getArea() float64 {
	return (0.5 * t.height * t.base)
}

func printArea(s shape) {
	fmt.Println("Area is", s.getArea())
}
