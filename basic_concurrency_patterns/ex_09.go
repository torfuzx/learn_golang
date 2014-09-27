package main

import (
	"fmt"
	"sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		// we ignore done here because these sends cannot block
		out <- n
	}
	close(out)
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // HL
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return // HL
			}
		}
		close(out)
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs. output copies values from
	// c to out until c or done is closed, then calls wg.Done
	output := func(c <-chan int) {
		defer wg.Done() // HL
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return // HL
			}
		}
	}
	// ... the rest is unchanged

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// start a goroutine to close out once all the output goroutine are done. This must
	// start after the wg.Add call
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// set up a done channel that's shared by the whole pipeline, and close that channel
	// when this pipeline exits, as a signal for all the goroutines we started to exit
	done := make(chan struct{}) // HL
	defer close(done)

	in := gen(done, 2, 3)

	// distribute the sq work across two goroutiens that both read from in.
	c1 := sq(done, in)
	c2 := sq(done, in)

	// consume the first value from output
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// done will be closed by the deferred call // HL
}
