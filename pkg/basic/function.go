package basic

import (
	"fmt"
	"golang/pkg/utils"
)

// package level variables, accessible in this file's package
var packageVar1 int = 1
var packageVar2 string

const packageVar3 int = 1
const packageVar4 = "s"

// outside a func, a statement should always start with a keyword
// fmt.Print("Hello, world") illegal

// param a and b are both int
// returns 1 int value and 1 string value
func F1(a, b int) (int, string) {
	fmt.Printf("calling basic.F3() \n")

	return a + b, "test"
}

// with single return
func F2() int {
	fmt.Printf("calling basic.F2() \n")

	utils.Test1()

	return 1
}

// not necessary to return a value,
// function without return value
func F3() {
	fmt.Printf("calling basic.F3() \n")
	mapTest()
}

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
	// print 1432
}
