package main

import (
	"fmt"
	"golang/pkg/basic"
	"golang/pkg/utils"
)

// The main() function must appear only once, be in the main() package, and receive and return no arguments.
func main() {
	basic.F3()
	utils.Test1()

	user := basic.User{ID: 1, Name: "ethan", Email: "test"}

	fmt.Println(user.Name)
}
