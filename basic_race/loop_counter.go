package main

import (
	"fmt"
	"sync"
)

// The vairable i in the function is the same variable used by the loop, so the
// read in the goroutine races with the loop increment.
func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // not the 'i' you are looking for
		}()
	}

	wg.Wait()
}
