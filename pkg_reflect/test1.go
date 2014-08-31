package main

import (
	"fmt"
	"reflect"
)

type S struct {
	// Please note that:
	// 1. The field and field value shouldn't be separated by any
	// space.
	// 2. The field value could be closely followd by the field name without
	// space in the middle.
	// 3. The field value is quoted with double-quotes.
	F string `name:"Kurt Zhong"email:"kurtzhong520@gmail.com"`
}

func main() {
	s := S{}
	st := reflect.TypeOf(s)
	fmt.Println(st)

	// returns the first field of struct
	field := st.Field(0)

	fmt.Println(field.Tag.Get("name"))
	fmt.Println(field.Tag.Get("email"))
}
