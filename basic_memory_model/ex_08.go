// An example on incorrect synchronization.
// Problem:
// There is no guarantee that in doprint, observing the write to done implies
// observing the write to a.
// This version can incorrecly print an empty string instead of "hello world"
package main

import (
	"log"
	"os"
	"sync"
)

var a string
var done bool
var logger *log.Logger
var once sync.Once
var wg sync.WaitGroup

func init() {
	// setup logger
	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Open log file failed:", err)
	}
	logger = log.New(f, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.Println("set up log...")
}

func setup() {
	a = "hello world"
	done = true
}

func doprint() {
	defer func() {
		logger.Println("doprintln exit")
	}()
	logger.Println("doprint do...")
	logger.Printf("doprint done = %v\n", done)
	if !done {
		once.Do(setup)
	}
	logger.Printf("doprint a = %v\n", a)
	wg.Done()
}

func twoprint() {
	defer func() {
		logger.Println("twoprint done")
	}()
	logger.Println("twoprint do...")

	wg.Add(2)
	go doprint()
	go doprint()
}

func main() {
	twoprint()
	wg.Wait()
}
