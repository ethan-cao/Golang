package basic

import "fmt"

func slice1() {
	// Slice has dynamic size
	// a slice is a dynamically-sized, flexible view into the elements of an array.
	// A slice does not store any data itself; it is simply a projection of an underlying array.

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

// Slicing
func TestClone1() {
	original := []Human{Human{ID: 1}, Human{ID: 2}, Human{ID: 3}}
	clone := original[:]

	// copied is created by slicing original from the first element to the end.
	// This does NOT create a new underlying array. Instead, copied shares the same underlying array as original.
	// Any changes to the elements of copied will be reflected in original because they are looking at the same array.
	// However, if you append to copied and it exceeds the capacity of the original slice, it will allocate a new array and will no longer share the same underlying array with original.

	// different address
	fmt.Printf("original: %p \n", &original)
	fmt.Printf("copid: %p \n", &clone)

	// same address
	fmt.Printf("original: %p \n", &original[0])
	fmt.Printf("copied: %p \n", &clone[0])

	original[0].ID = 4
	fmt.Println(original[0].ID)
	fmt.Println(clone[0].ID)

}

// Direct assignment
func TestClone2() {
	original := []string{"Apple", "Banana", "Cherry"}

	clone := original

	fmt.Printf("copied: %p \n", &clone[0])
}

// DeepClone
func TestClone3() {
	original := []Human{Human{ID: 1}, Human{ID: 2}, Human{ID: 3}}
	clone := make([]Human, len(original))
	copy(clone, original)

	// you are creating a new slice copied with the same length as original
	// then copy the elements from original into copied. This is a shallow copy.
	// The copied slice is a different slice that has the same elements as original,
	// but it has its own underlying array. Changes to the elements of copied will not affect original, and vice versa.

	// different address
	fmt.Printf("original: %p \n", &original)
	fmt.Printf("copid: %p \n", &clone)

	// different address
	fmt.Printf("original: %p \n", &original[0])
	fmt.Printf("copied: %p \n", &clone[0])

	original[0].ID = 4
	fmt.Println(original[0].ID)
	fmt.Println(clone[0].ID)
}

func TestClone4() {
	original := []string{"Apple", "Banana", "Cherry"}

	var clone []string
	clone = append(clone, original...)

	fmt.Printf("copied: %p \n", &clone[0])
}
