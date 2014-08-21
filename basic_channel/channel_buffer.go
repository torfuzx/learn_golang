/*!

Channel buffering
-----------------

By default, channels are unbuffered, meaning that they will only accept send
(chan <-) if there is a corresponding receiver ( <- chan ) ready to receive the 
to-be-sent value. 

Buffered channels accept a limited number of values witout a corresponding 
receiver for those values.

*/



package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 4)
	messages <- "buffered"
	messages <- "channel"
	messages <- "is"
	messages <- "awesome"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
