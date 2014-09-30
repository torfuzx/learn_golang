package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0

	// simulate read
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// in order to ensure that this gorotine doesn't starve the sheduler,
				// we explicitly yield after each operation with runtime.Gosched(), this
				// yielding is handled atomatically with e.g. every channel operation
				// and for blocking calls like time.Sleep, but in this case we need to
				// do it manually
				runtime.Gosched()
			}
		}()
	}

	// simulate write
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(10)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
