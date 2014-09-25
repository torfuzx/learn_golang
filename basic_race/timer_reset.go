// Run this example using:
//	go run -race timer_reset.go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// -----------------------------------------------------------------------------
// Shows an unsynchronized read and write of the variable t from different
// goroutines. If the initial timer duration is very small, the timer function
// may fire before the main goroutine has assigned a value to t and so the call
// to t.Reset is made with a nil t.
// -----------------------------------------------------------------------------
func main() {
	start := time.Now()
	var t *time.Timer

	// time.AfterFunc waits for the duration to elapse and then calls f
	// in its own goroutine. It returns a Timer that can be used to cancel
	// the call using its stop method.
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		// timer.Reset() changes the timer to expire after duration d.
		// It returns true if the timer had been active, false if the timer
		// had expired or been stopped.
		t.Reset(randomDuration())
	})

	time.Sleep(3 * time.Second)
}

func randomDuration() time.Duration {
	// Nanosecond           1
	// Microsecond          1,000
	// Millisecond          1,000,000
	// Second               1,000,000,00
	return time.Duration(rand.Int63n(1e2))
}
