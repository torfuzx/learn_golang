package main

import "fmt"

var c = make(chan int, 10)
var a string

func f() {
	a = "hello world"
	c <- 0
}

func main() {
	go f()
	<-c
	fmt.Println(a)
}
