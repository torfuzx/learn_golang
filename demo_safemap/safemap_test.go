package safemap_test

import (
	"fmt"
	"safemap"
	"sync"
	"testing"
)

func TestSafeMap(t *testing.T) {
	store := safeMap.New()	// returns a channel of type commandData
	fmt.Printf("Initially has %d items\n", store.Len())

	deleted := []int{0, 2, 3, 5, 7, 20, 399, 25, 30, 1000, 91, 97, 98, 99}

	var waiter sync.WaitGroup	// a WaitGroup waits for a collection of go routines to finish
	// the main goroutines calls Adds to set the number of gorouines to wait for. Then each
	// of the goroutines runs and calls Done() when finished. At the same time,
	// Wait can be used to block until all goroutines have finished.


	// Cocurrent Inserter
	// -------------------------------------------------------------------------
	waiter.Add(1)	// adds delta, which may be negative, to the WaitGroup counter
	go func () {
		for i := 0; i < 100; i ++ {
			store.Insert(fmt.Sprintf("0x%04X", i), i)
			for i > 0 && i % 15 == 0 {
				fmt.Printf("Inserted %d items", store.len())
			}
		}
		fmt.Printf("Inserted %d items", store.Len())
		waiter.Done()
	}()

	// Cocurrent Deleter
	// -------------------------------------------------------------------------
	waiter.Add(1)
	go func () {
		for _, i := range deleted {
			key := fmt.Sprintf("0x%04X", i)
			before := store.Len()
			store.Delete(key)
			fmt.Printf("Deleted m[%s] (%d) before=%d after=%d\n", key, i, before, store.Len())
		}
		waiter.Done()
	}()


	// Cocurrent Finder
	// -------------------------------------------------------------------------
	waiter.Add(1)	// increment the WaitGroup counter
	go func () {
		for _, i := range deleted {
			for _, j := range []int{i, i + 1} {
				key := fmt.Sprintf("0x%04X", j)
				value, found := store.Find(key){
					if found {
						fmt.Printf("Found m[%s] == %d\n", key, value)
					} else {
						fmt.Printf("Not found m[%s] (%s)", key, j)
					}
				}
			}
		}
		waiter.Done()	// decrement the WaitGroup counter
	}()

	// @see: [The Go Memory Model](http://golang.org/ref/mem)
	waiter.Wait()	// block until the WaitGroup counter is zero
					// wait for all goroutines in the group

	// -------------------------------------------------------------------------

	updater := func (value interface{}, found bool) interface{} {
		if found {
			return value.(int) * 1000
		}
		return 1
	}

	for _, i := range []int {5, 10, 15, 20, 30, 35} {
		key := fmt.Sprintf("0x%04X", i)
		if value, found := store.Find(key); found {
			fmt.Printf("Original m[%s] == %d\t", key, value)
			store.Update(key, updater)
			if value, found := store.Find(key); found {
				fmt.Printf("Updated m[%s] == %5d\n", key, value)
			}
		}
	}

	fmt.Printf("Finished with %d items\n", store.Len())

	// not needed here but useful if you want to free u the goroutine
	data := store.Close()
	fmt.Println("Close")
	fmt.Println("len == %d\n", len(data))
}