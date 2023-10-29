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

}
