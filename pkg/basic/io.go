package basic

import "fmt"

func PrintFormat() {
	date := "Monday"
	weather := "sunny"
	fmt.Printf("today is %v.\nThe weather is %v", date, weather)

	name := "ethan"
	fmt.Printf("name is of type %T", name)
}

func ReadFromCLI() {
	var userName string

	fmt.Scan(&userName)
}
