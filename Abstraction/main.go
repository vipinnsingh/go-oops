package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	// Using abstraction to handle different shapes
	printArea(c)
	printArea(r)
}

/*
	Here, Shape is an abstraction, and Circle and Rectangle are concrete implementations.
	The printArea function works with any type that satisfies the Shape interface without
	needing to know the specific type.
*/
