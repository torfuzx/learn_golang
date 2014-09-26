package main

import "fmt"

var a string

func hello() {
	go func() {
		a = "hello"
	}()
	fmt.Printf("a = [%v]\n", a)
}

func main() {
	hello()
}
