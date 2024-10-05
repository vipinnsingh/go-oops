package main

import "fmt"

type Engine struct {
	Horsepower int
}

func (e *Engine) Start() {
	fmt.Println("engine running")
}

type Wheel struct {
	Dimension int
}

type Car struct {
	CarName string
	Engine
	Wheel
}

func NewCar(carName string) Car {

	car := Car{
		CarName: carName,
		Engine: Engine{
			Horsepower: 7000,
		},
		Wheel: Wheel{
			Dimension: 16,
		},
	}

	return car

}

func main() {

	car := NewCar("City")

	fmt.Printf("car: %v\n", car)

	car.Start()

}
