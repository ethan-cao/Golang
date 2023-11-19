package basic

import "fmt"

func slice1() {
	// Slice has dynamic size

	var s1 []string
	var s2 = []string{}
	s3 := []string{}
	s4 := []string{"a", "b"}

	s1 = append(s1, "a")
	s1 = append(s1, "b")
	s1 = append(s1, "c")

	fmt.Printf("size %v \n", len(s1)) // 3
	fmt.Printf("size %v \n", len(s2)) // 0
	fmt.Printf("size %v \n", len(s3)) // 0
	fmt.Printf("size %v \n", len(s4)) // 0

	// empty slice of length 3
	s5 := make([]string, 3)
	fmt.Printf("size %v \n", len(s5)) // 3

	for idx, value := range s1 {
		fmt.Printf("value from index %v is %v \n", idx, value)
	}

	// arrays and slices cannot be declared as constants because their values are not compile-time constants
}

func TestCopy1() {
	original := []User{User{ID: 1}, User{ID: 2}, User{ID: 3}}
	copid := original[:]

	// copied is created by slicing original from the first element to the end.
	// This does not create a new underlying array. Instead, copied shares the same underlying array as original.
	// Any changes to the elements of copied will be reflected in original because they are looking at the same array.
	// However, if you append to copied and it exceeds the capacity of the original slice, it will allocate a new array and will no longer share the same underlying array with original.

	// different address
	fmt.Printf("original: %p \n", &original)
	fmt.Printf("copid: %p \n", &copid)

	// same address
	fmt.Printf("original: %p \n", &original[0])
	fmt.Printf("copied: %p \n", &copid[0])

	original[0].ID = 4
	fmt.Println(original[0].ID)
	fmt.Println(copid[0].ID)

}

func TestCopy2() {
	original := []User{User{ID: 1}, User{ID: 2}, User{ID: 3}}
	copied := make([]User, len(original))
	copy(copied, original)

	// you are creating a new slice copied with the same length as original
	// then copy the elements from original into copied. This is a shallow copy.
	// The copied slice is a different slice that has the same elements as original,
	// but it has its own underlying array. Changes to the elements of copied will not affect original, and vice versa.

	// different address
	fmt.Printf("original: %p \n", &original)
	fmt.Printf("copid: %p \n", &copied)

	// different address
	fmt.Printf("original: %p \n", &original[0])
	fmt.Printf("copied: %p \n", &copied[0])

	original[0].ID = 4
	fmt.Println(original[0].ID)
	fmt.Println(copied[0].ID)
}
