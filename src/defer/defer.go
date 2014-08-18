/**

defer
-----
- A defer statement pushes a function onto a list.
- The list of saved call is executed after the surrounding functions returns.
- Defer is commonly used to simplify functions that perform various clean-up
  actions.

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
