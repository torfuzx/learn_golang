// There is a pattern to our pipeline functions:
// - stages close their outbound channels when all the send operations are done
// - stages keep receiving value from inbound channels until those channels are closed

// This pattern allows each receiving stage to be written as a range loop and ensures that
// all goroutines exit once all values have been successfully sent downstream

// But in real pipelines, stages don't always receive all the inbound values. Sometimes
// this is by design: the receiver may only need a subset of values to make progress. More
// often, a stage exits early because an inbound value represents an error in an earlier
// stage. In either case the receiver should not have to wait for the remaining value to
// arrive, and we want earilier stages to stop producing values that later stages don't
// need.
package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// start an output goroutine for each input channel in cs. output copies value from
	// c to out until c is closed, then calles wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	// start a goroutine to close out once all the output goroutine are done. This must
	// start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := gen(2, 3)

	// distribute the sq work across two goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	// consume the merged output from c1 and c2
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
