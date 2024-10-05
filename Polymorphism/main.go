package main

import "fmt"

type Shape interface {
	Render()
}

type Circle struct {
	Radius int
}

func (c *Circle) Render() {
	fmt.Println(c.Radius)
}

type Square struct {
	Length int
}

func (s *Square) Render() {
	fmt.Println(s.Length)
}

func main() {

	var x Shape = &Circle{
		Radius: 9,
	}
	x.Render()

	var y Shape = &Square{
		Length: 5,
	}
	y.Render()
}

/*
	Polymorphism in OOP allows objects of different types to be treated as instances of the same type.
	Go achieves this via interfaces, which define a set of methods that any type can implement.
	In Go, if a type implements all the methods of an interface, it is automatically considered to satisfy that interface.
	This enables polymorphic behavior.
*/
