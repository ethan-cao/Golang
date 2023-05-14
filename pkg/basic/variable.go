package basic

import "fmt"

func varTest() {
	// cannot re-declare, mutatable
	var name string = "John"

	// declare multiple, type inferred
	var b, c = 1, 2

	// short variable declaration(declare, infer tpye, assign), mutable
	// same as var foo int = 32
	foo := 32

	// wrong, foo is of type int
	// foo = "t"

	// nil is a special value that can be assigned to any data type.

	// immutable variable, type is inferred
	const pi = 3.14159

	// declare multiple constants in a single statement,
	const (
		v1 = 10
		v2 = 20
		v3 = 30
	)

	// "" for string
	string1 := "this is a string"
	// '' for character
	char1 := 'c'

	fmt.Println(name, b, c, foo, string1, char1)
}

func PointerTest1() {
	var num int = 1

	// use & to get the memory address
	fmt.Printf("memory address of num is %v \n", &num)

	// store the address of int num in pointer type address
	var address *int = &num
	fmt.Printf("memory address of num is %v \n", address)

	// use * to get the value
	fmt.Printf("value stored in this memory address is %v \n", *address)

	// de-reference a pointer to access the value
	// update value store in address to 2, so num is updated to 2
	*address = 2
	fmt.Println(num) // 2

	name := "name1"
	// pass-by-value: parameter is copied into function
	update1(&name)
	fmt.Println("after update1(): ", name)

	slice := []int{}
	update2(&slice)
	fmt.Printf("%v", slice)
}

func update1(name *string) {
	// de-reference a pointer to access the value
	*name = "name2"
}

func update2(slice *[]int) {
	*slice = append(*slice, 1)
	*slice = append(*slice, 2)
	*slice = append(*slice, 3)
}

func zeroTest() {
	/*
		zero value is the default value assigned to a variable when it is declared but not explicitly initialized.
			Numeric types (int, float, complex): 0
			Boolean: false
			Strings: ""
			Pointers: nil
			Interfaces: nil
			Slices: nil
			Channels: nil
			Maps: nil
			Functions: nil
			Structs: The zero value of each field in the struct
	*/
}
