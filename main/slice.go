package main

import "fmt"

func slice1() {
	// Slice has dynamic size

	// empty slice of length 0
	var s1 []string
	var s2 = []string{}
	s3 := []string{}

	s1 = append(s1, "a")
	s1 = append(s1, "b")
	s1 = append(s1, "c")

	fmt.Printf("size %v \n", len(s1)) // 3

	// empty slice of length 3
	s4 := make([]string, 3)

	for idx, value := range s1 {
		fmt.Printf("value from index %v is %v \n", idx, value)
	}
}
