package main

import (
	"fmt"
	"time"
)

var timeout chan bool
var result chan int

func init() {
	// the timeout channel is buffered iwth space for 1 value, allowing the timeout
	// goroutine to send to the channel and then exit.
	timeout = make(chan bool, 1)
	result = make(chan int)
}

func main() {
	go func() {
		var i = 0
		// do hard work, takes too much time
		time.Sleep(time.Second * 2)
		result <- i
	}()

	// if nothing arrives on result channel within 1 one second, the timeout case is
	// selected and the attempt to read from result channel is abandomed
	select {
	case n := <-result:
		// a read from ch has occured
		fmt.Println("Received ", n)

	case <-time.After(1 * time.Second):
		// the read from ch has timed out
		fmt.Println("1 second elapsed, wait no more.")
	}
}
