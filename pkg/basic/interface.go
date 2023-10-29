package basic

import "fmt"

// Interfaces are named collections of method signatures.

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width, height float64
}

// implement the area()
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// implement the perimeter()
func (r Rectangle) perimeter() float64 {
	return r.width*2 + r.height*2
}

func interfaceTest(
	g Geometry,
	any interface{}, // a special type that can hold any value, any type
) {
	g.area()

	rect := Rectangle{width: 5, height: 3}
	var geometry Geometry = rect
	fmt.Println(geometry.area())
}
