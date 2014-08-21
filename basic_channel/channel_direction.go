/*!

We can specify whether a channel is used for receivin or sending.

*/

package main

import "fmt"

// ping only accept a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong accept first a channel for receving value
// and then a channel for sending value
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
