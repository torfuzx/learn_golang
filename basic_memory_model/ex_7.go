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

\

func main() {
	go f()
	g()
	<-done
}
