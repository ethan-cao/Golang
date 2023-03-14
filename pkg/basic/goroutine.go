package basic

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

func RunLongTask() {
	// set waitGroup counter to be 2 before spin off goroutine
	// so the main thread waits for these thread to complete
	waitGroup.Add(2)

	// spin off new goroutine to run tasks
	go runLongTask1()
	go runLongTask2()

	fmt.Printf("main thread task finshed\n")

	// wait until waitGroup counter to be 0
	waitGroup.Wait()
}

func runLongTask1() {
	time.Sleep(5 * time.Second)

	fmt.Printf("long task1 finshed\n")

	// reduce waitGroup counter by 1
	waitGroup.Done()
}

func runLongTask2() {
	time.Sleep(10 * time.Second)

	fmt.Printf("long task2 finshed\n")

	// reduce waitGroup counter by 1
	waitGroup.Done()
}
