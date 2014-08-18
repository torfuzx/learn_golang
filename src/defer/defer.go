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

panic
-----

- Stops the ordinary flow and begins paniking. When the function f calls panic,
  execution of f stops, any defferred function in F are executed normally, and
  then f returns to its caller. F then behaves likes a call to panic. The
  process continues up the stack until all functions in the current goroutine
  have returned, at which point the program crashes. Panic can be initiated by
  invoking panic directly. They can also be caused by runtime errors, such as
  out-of-bound array access.

recover
-------

- A built-in function that regains controls of paniking goroutine.
  Only useful inside deferred functions.
- During normal executions, a call to recover will return nil and have no other
  effects.
- If the current goroutine is paniking, a call to recover will capture
  the value given to panic and resume normal execution.


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

	// test panic and defer
	println("-----------------------------------------------------------------")
	f2()
	fmt.Println("Returned normally from f.")
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

func f2() {
	// if we remove this defer statement, the panic is not recovered and reaches
	// the top of the goroutine's call stack,terminating the program.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
