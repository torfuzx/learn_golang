package main

import (
	"fmt"
	"reflect"
)

type Dish struct {
	Id     int
	Name   string
	Origin string
	Query  func()
}


func attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)

	// if a pointer to a struct is passed, get the type of thee dereferenced
	// object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// create an attribute data structuree as a map of types keyed by a string
	attrs := make(map[string]reflect.Type)

	// only structs are supported so return a empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// iterate throught the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}
	return attrs
}

func main() {
	// iterate through the attributes of a Data Model instance
	for name, mtype := range attributes(&Dish{}) {
		fmt.Printf("Name: %s, Type: %s\n", name, mtype.Name())
	}
}
