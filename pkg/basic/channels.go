package basic

import "fmt"

// send values into channels from one goroutine and receive those values into another goroutine.

func ChannelTest() {
	// define a channel
	channel1 := make(chan string)

	// start a new goroutine to send a value to the channel
	go func() {
		// send a value to the channel
		channel1 <- "ping"
	}()

	message := <-channel1
	fmt.Printf("message inside channel: %s\n", message)
}
