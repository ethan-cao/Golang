package main

// Interfaces are named collections of method signatures.

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return r.width*2 + r.height*2
}

func interfaceTest(
	g geometry,
	any interface{}, // a special type that can hold any value, any type
) {
	g.area()
}
