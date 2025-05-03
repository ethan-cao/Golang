package main

import (
	"fmt"
)

func main() {
	fmt.Print("Go runs on ")

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

	if val, ok := m1["k1"]; ok {
		fmt.Printf("@@@ %v\n", val)
	}

	v, ok := m1["k1"]
	fmt.Printf("v: %v, ok: %v\n", v, ok)
}
