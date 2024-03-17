package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// There is ONLY one main() in the main package
// it's the entry point for all Go application, this file has to be called main with func main()
// it receives and returns no arguments
func main() {

	err := testFunc()

	if err != nil {
		fmt.Printf("An error occurred: %+v\n", err)
		os.Exit(1)
	}
	fmt.Println("Task completed successfully")

}

func testFunc() error {
	return errors.WithStack(errors.New("with error"))
	// return errors.New("error")
}
