package main

import (
	"log"
	"os"
	"sync"
)

var a string
var once sync.Once
var logger *log.Logger
var wg sync.WaitGroup

func setup() {
	logger.Println("setup do...")
	a = "hello world"
}

func doprint() {
	defer func() {
		logger.Println("doprint done")
	}()
	logger.Println("doprint do...")
	// Do calls the funtion if and only if Do is being called for the first time
	// for this instance of Once.
	once.Do(setup)
	logger.Println(a)
	wg.Done()
}

func twoprint() {
	wg.Add(2)
	go doprint()
	go doprint()
	wg.Wait()
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
	twoprint()
}
