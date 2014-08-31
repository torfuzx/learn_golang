/*******************************************************************************

Example originated from the go reflect package officical documentation.

*******************************************************************************/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// swap is the implmentation passed to MakeFunc. It must work in terms of
	// reflect.Value so that it is possible to write code without knowing
	// beforehand what the type will be.
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap expects fptr to be a pointer to a nil funciton. It sets that
	// pointer to a new function created with MakeFunc. When the function is
	// invoked, reflect turns the arguments into values, calls swap, and then
	// turns swap's result slice into the values returned by the new function.
	makeSwap := func(fptr interface{}) {
		// fptr is a pointerr to a function. Obtain the function value itself,
		// (like nil) as a reflect.Value, so that we can query its type and then
		// set the value.
		fn := reflect.ValueOf(fptr).Elem()

		// make a fucntion of the right type
		v := reflect.MakeFunc(fn.Type(), swap)

		// assign it to the value fn represents
		fn.Set(v)
	}

	// make and call a swap function for ints
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// make and call a swap function for float64s
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
}
