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
