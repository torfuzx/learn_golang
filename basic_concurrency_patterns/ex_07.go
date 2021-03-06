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

func merge(done <-chan struct{}, cs ...chan<- int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// start an output goroutine for each input channel in cs. output copies value from
	// c to out until c is closed, then calles wg.Done
	output := func(c <-chan int) {
		for n := range c {
			select {

			case out <- n:
			case <-done:
			}
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

	// distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// consume the first value from output
	done := make(chan struct{}, 2)
	out := merge(done, c1, c2)
	fmt.Println(<-out)

	// tell the remaining senders we're leaving
	done <- struct{}{}
	done <- struct{}{}
}
