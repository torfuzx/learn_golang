package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var f func()
	f = func() {
		fmt.Println(time.Now().Sub(start))
		time.AfterFunc(time.Duration(rand.Int63(1e9)), f)
	}
	time.AfterFunc(time.Duration(rand.Int63n(1e9)), f)
	time.Sleep(5 * time.Second)
}
