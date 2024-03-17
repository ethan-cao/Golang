package basic

import (
	"fmt"
	"sync"
	"time"
)

// A goroutine is a lightweight thread of execution.

var waitGroup = sync.WaitGroup{}

func RunLongTask() {
	// set waitGroup counter to be 2 before spin off goroutine
	// so the main thread waits for these thread to complete, when calling waitGroup.Wait()
	waitGroup.Add(2)

	// call the function using go keyword, whcih start a goroutine running concurrently
	// spin off new goroutine to run tasks
	go runLongTask1()
	go runLongTask2()

	fmt.Printf("main thread task finished\n")

	// wait until waitGroup counter to be 0
	waitGroup.Wait()
}

func runLongTask1() {
	time.Sleep(5 * time.Second)

	fmt.Printf("long task1 finshed in 5s \n")

	// reduce waitGroup counter by 1
	waitGroup.Done()
}

func runLongTask2() {
	time.Sleep(10 * time.Second)

	fmt.Printf("long task2 finshed in 2s\n")

	// reduce waitGroup counter by 1
	waitGroup.Done()
}
