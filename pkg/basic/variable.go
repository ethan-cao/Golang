package basic

import "fmt"

func varTest() {
	// cannot re-declare, mutable
	var name0 string = "Tom"
	// infer string type
	var name1 = "Hanks"
	// infer string type
	name2 := "New"

	// wrong, name2 is of type string
	// name2 = 1

	// declare multiple variable, type inferred
	var b, c = 1, 2

	// nil is a special value that can be assigned to any data type.

	//  cannot re-declare, immutable
	const constantVariable int = 1

	// type is inferred
	const pi = 3.14159

	// wrong, const cannnot use :=
	// const num1 := 1

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

	fmt.Println(name0, name1, name2, b, c, string1, char1)
}

func PointerTest1() {
	// var value int = 1
	value := 1

	// use & to get the memory address
	fmt.Printf("memory address of num is %v \n", &value)

	// store the address of int value in pointer type address
	var address *int = &value
	fmt.Printf("memory address of num is %v \n", address)

	// use * to get the value: de-reference
	fmt.Printf("value stored in this memory address is %v \n", *address)

	// update value store in address to 2, so value is updated to 2
	*address = 2
	fmt.Println(value) // 2

	name := "name1"
	fmt.Printf("memory address of name is %v \n", &name)

	// pass-by-value: parameter is copied into function
	update0(name)
	fmt.Println("after update0(): ", name) // name1, unchanged
	update1(&name)
	fmt.Println("after update1(): ", name) // name2

	slice := make([]int, 3)
	update2(slice)
	fmt.Printf("after update2() %v \n", slice) // [], unchanged
	update3(&slice)
	fmt.Printf("after update3() %v \n", slice) // [1,2,3]
}

func update0(name string) {
	fmt.Printf("memory address of name param is %v \n", &name)
	// pass-by-value: value "name1" is copied to function param name,
	// which has a different memory address than the name variable from caller
	// any change made is made only to the param name
	name = "name2"
}

func update1(name *string) {
	fmt.Printf("memory address of name param is %v \n", &name)
	// pass-by-value: address of the argument in caller is copied to function param name
	// any change made is made to the value stored at the argument's address
	*name = "name2"
}

func update2(slice []int) {
	slice[0] = 9
	slice[1] = 9
	slice[2] = 9
	// slice = append(slice, 1)
	// slice = append(slice, 2)
	// slice = append(slice, 3)
}

func update3(slice *[]int) {
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
