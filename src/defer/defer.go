/**

defer
-----
- A defer statement pushes a function onto a list.
- The list of saved call is executed after the surrounding functions returns.
- Defer is commonly used to simplify functions that perform various clean-up
  actions.

Three simple rules of defer:
- 1. A deferred fucntion's arguments are evaluated when the defer statement is
  evaluated.
- 2. Deferred function calls are executed in Last In First Out order after the
  surrounding function returns.
- 3. Deferred funcitons may read and assign to the returning funciton's named
  return values.


*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	res := f()
	fmt.Println(res)

	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}

	deferRule1()
	deferRule2()
	i := deferRule3()
	fmt.Println("value:", i)
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

// A defer example to shows what defer want to solve
func CopyFileWithoutDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Open(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

func CopyFileWithDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Open(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func deferRule1() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func deferRule2() {
	for i := 0; i <= 4; i++ {
		defer fmt.Println(i)
	}
	return
}

func deferRule3() (i int) {
	defer func() { i++ }()
	return 1
}
