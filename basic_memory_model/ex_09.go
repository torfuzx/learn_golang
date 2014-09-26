// An example on incorrect synchronization
// 1. There is no guarantee that observing the write to `done` implies write
// to `a`, so this program could print an empty string too.
// 2. Worse, there is no guarantee that the write to `done` will ever be
// observed by main(),since there are no synchronization event between the two
// threads. The loop in main is not guaranteed to finish.
package main

import (
	"log"
	"os"
)

var a string
var done bool
var logger *log.Logger

func setup() {
	a = "hello world"
	done = true
}

func init() {
	// setup logger
	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Open log file failed:", err)
	}
	logger = log.New(f, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.Println("set up log...")
}

func main() {
	go setup()
	for !done {
	}
	logger.Println("a = %v\n", a)
}
