package main

import (
	"fmt"
)

// Struct example
type Rectangle struct {
	width  float64
	height float64
}

// Method for Rectangle struct
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Another struct
type Circle struct {
	radius float64
}

// Method for Circle struct
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Interface example
type Shape interface {
	Area() float64
}

func interfaceMain() {
	// Using structs
	rect := Rectangle{
		width:  10,
		height: 5,
	}
	fmt.Printf("Rectangle area: %.2f\n", rect.Area())

	circle := Circle{
		radius: 5,
	}
	fmt.Printf("Circle area: %.2f\n", circle.Area())

	// Using interface
	// Both Rectangle and Circle automatically implement Shape interface
	var shape Shape
	shape = rect
	fmt.Printf("Shape (Rectangle) area: %.2f\n", shape.Area())

	shape = circle
	fmt.Printf("Shape (Circle) area: %.2f\n", shape.Area())

	// Demonstrating interface polymorphism
	shapes := []Shape{rect, circle}
	for _, s := range shapes {
		fmt.Printf("Shape area: %.2f\n", s.Area())
	}
}
