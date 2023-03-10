package main

import "fmt"

// outside a func, a statement should always start with a keyword
// fmt.Print("Hello, world") illegal

// param a and b are both int
// it returns 1 int value and 1 stirng value
func f1(a, b int) (int, string) {
	fmt.Print("Hello, world")

	return a + b, "test"
}

// with single return
func f2() int { return 1 }

// not necessry to reutrn a value,
// function without reutrn value
func f3() { fmt.Print("Hello, world") }

// Named return values
func f5(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// A return statement without arguments returns the named return values.
	// This is known as a "naked" return.
	return
}

func DeferTest() {
	fmt.Print("1")

	// defers the execution at the end of the enclosing function
	// defer push the call to a stack
	defer fmt.Print("2")
	defer fmt.Print("3")

	fmt.Print("4")
}

// print 1432
