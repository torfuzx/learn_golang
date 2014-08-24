package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// run it synchronously
	f("direct")

	// invoke in a goroutine, will be execute cocurrently with the calling
	// goroutine
	go f("go routine")

	// start goroutine for anoymous funciton call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string

	fmt.Scanln(&input)
	fmt.Println("done")
}
