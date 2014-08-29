package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F string `name:"Kurt Zhong" email: "kurtzhong520@gmail.com"`
	}

	s := S{}
	//
	st := reflect.TypeOf(s)
	fmt.Println(st)

	// returns the first field of struct
	field := st.Field(0)

	fmt.Println(field.Tag.Get("emailz"))
}
