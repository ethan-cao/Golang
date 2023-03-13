package basic

import "fmt"

func mapTest() {
	// create an empty map string->int
	m0 := make(map[string]int)
	fmt.Println(len(m0))

	// create a map string->int with initial values
	m1 := map[string]int{
		"k1": 1,
		"k2": 2,
	}

	m1["k3"] = 3

	fmt.Println(len(m1)) // 3

	for k, v := range m1 {
		fmt.Println("Key:", k, "Value:", v)
	}
}
