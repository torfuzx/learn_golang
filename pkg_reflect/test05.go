package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	I int
}

func main() {
	// a random struct
	v := MyStruct{10}
	// an interface{} containingg concrete value v
	var iv interface{} = v

	fmt.Println(reflect.ValueOf(v).Kind())
	fmt.Println(reflect.ValueOf(&v).Elem().Kind())

	fmt.Println()

	val := reflect.ValueOf(&iv)
	fmt.Println(val.Type(), ",", val.Kind())
	fmt.Println(val.Elem().Type(), ",", val.Elem().Kind())
	fmt.Println(val.Elem().Elem().Type(), ",", val.Elem().Elem().Kind())

	fmt.Println()

	fmt.Println(val.Elem().Elem().Field(0).Int())

	fmt.Println()

	// assign a string to what was a struct (just as valid ass iv ="string")
	fmt.Println(val.Elem().Kind())
	val.Elem().Set(reflect.ValueOf("a string  which is obviously not a struct"))
	fmt.Println(val.Elem().Elem().String())
}
