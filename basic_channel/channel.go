/*!

Channels
--------

Channels are the pipes that connect concurrent goroutines. You can send value
into channels from goroutine and receive those values from another goroutine.

- send a value into channel: channel <- 
- recv a value from channel: <- channel 

Just imagagine the channel is queue, we push a message at it's end, 
and then we pop a value from the queue's head

By default, send and receive blocks until both the sender and receiver are 
ready. The property allows us to wait a for the message without having to use 
any other synchronization.

*/

package main

import (
	"fmt"
)

func main() {
	// create a channel by using the syntax:
	// make(chan VALUE_TYPE)
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages

	fmt.Println(msg)
}
