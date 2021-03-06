// Some random package doc that isn't in the right form.
// MATCH /package comment should.*form.*"Package testdata .*"/
package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   int
	height int
}

type square struct {
	side int
}

func main() {
	t := triangle{4, 4}
	printArea(t)

	s := square{2}
	printArea(s)
}

func printArea(s shape) {
	fmt.Println("Area is: ", s.getArea())
}

func (t triangle) getArea() float64 {
	return (float64)(t.base * t.height / 2)
}

func (s square) getArea() float64 {
	return (float64)(s.side * s.side)
}
