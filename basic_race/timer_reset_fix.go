package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The main goroutine is wholly responsible for setting and resetting the Timer
// t and a new reset channel communicates the need to reset the timer in a
// thread-safe way.
func main() {
	start := time.Now()
	reset := make(chan bool)
	var t *time.Timer

	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		reset <- true
	})

	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration)
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
