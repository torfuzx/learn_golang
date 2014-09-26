// This is an example on incorect synchronization.
package main

import (
	"log"
	"os"
)

type T struct {
	msg string
}

var g *T
var logger *log.Logger

func setup() {
	defer func() {
		logger.Println("setup goroutine exit")
	}()
	logger.Println("setup goroutine doing...")
	t := new(T)
	t.msg = "hello world"
	g = t
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
	logger.Println("go statement done")
	// even if main observes g != nil and exits it loop, there is no guarantee
	// that it will observe the initialized value for g.msg
	for g == nil {
		// todo: why after i comment the following line, the program blocks
		logger.Println("looping...")
	}
	logger.Printf("g.msg = %v\n", g.msg)
}
