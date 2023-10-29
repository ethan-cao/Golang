package basic

import "fmt"

func array1() {
	// array is fixed-sized

	// define an empty array
	var a [5]int
	a[0] = 1

	b := [5]int{1, 2, 3} // [1 2 3 0 0]
	b[1] = 3             // [1 3 3 0 0]

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
