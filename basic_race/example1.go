package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	m := make(map[string]string)

	go func() {
		m["1"] = "a" // first conflicting access
		c <- true
	}()

	m["2"] = "b" // second conflicting access
	<-c

	for k, v := range m {
		fmt.Println(k, v)
	}
}
