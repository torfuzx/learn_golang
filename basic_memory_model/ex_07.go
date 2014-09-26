// An example on incorrect synchronization
// A read r may observes the value written by a write w that happens
// concurrently with r. Even if this occurs, it does not imply that reads
// happening after r will observe writes that happened before w.

package main

import (
	"log"
	"os"
)

var a, b int
var logger *log.Logger
var done chan bool

func f() {
	defer func() {
		logger.Println("f done")
	}()
	logger.Println("f do...")
	a = 1
	b = 2
	done <- true
}

func g() {
	defer func() {
		logger.Println("g done")
	}()
	logger.Println("g do...")
	logger.Printf("b = %v\n", b)
	logger.Printf("a = %v\n", a)
}

func init() {
	// setup logger
	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Open log file failed:", err)
	}
	logger = log.New(f, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.Println("set up log...")

	// set up the channel done
	done = make(chan bool)
}

func main() {
	go f()
	g()
	<-done
}
