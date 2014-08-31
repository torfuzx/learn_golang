/*******************************************************************************

- ValueOf() returns a new value initialized to the concrete value stored in the
interface i.
- Value.FieldByName() returns the struct field with the given name
- Value.Int() returns underlying vlaue, as a int64
- Value.Elem() returns the value that the interface v contains or that the
pointer v points to.

*******************************************************************************/

package main

import (
	"fmt"
	"reflect"
)

type S struct {
	// exported field
	N int
}

func main() {
	s := S{1}

	// get
	immutable := reflect.ValueOf(s)
	val := immutable.FieldByName("N").Int()
	fmt.Printf("n=%d\n", val)

	// set
	mutable := reflect.ValueOf(&s).Elem()
	mutable.FieldByName("N").SetInt(7)
	fmt.Printf("n=%d\n", s.N)
}
