/*!

Channel synchronization
-----------------------

We can use channels to synchronize execution across gotoutines. 
 


*/

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	// the done channel is used to notify that the working goroutine has 
	// finished its work
	done := make(chan bool, 1)
	
	// start a goroutine, giving it a channel to notify on
	go worker(done)

	// use a blocking receive to wait for a goroutine to finish
	// block until we receive a notification from the worker on the channel
	<-done
}
