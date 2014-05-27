package main

import (
	"fmt"
)

func main() {
	res := f();
	fmt.Println(res);

	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func f() (result int) {
	defer func () {
		result ++
	}()
	return 0
}
