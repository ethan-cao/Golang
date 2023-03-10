package main

import "fmt"

func array1() {
	// An array in Go is a fixed-sized

	// define an empty array
	var a [5]int
	a[0] = 1

	b := [5]int{1, 2, 3}
	b[1] = 3

	fmt.Printf("size of array %v \n", len(b)) // 5

	for idx, value := range b {
		fmt.Printf("value from index %v is %v \n", idx, value)
	}

	var twoDArray [][]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDArray[i][j] = i + j
		}
	}

}
