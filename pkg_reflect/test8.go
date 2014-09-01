/*******************************************************************************

Law of reflection

http://blog.golang.org/laws-of-reflection



Types and interface
-------------------

Go is statically typed. Every variable has a static type, that is exactly one
type known and fixed at compile time.

One important category of type is interface types, which represent fixed set of
method. An interface variable can store any concrete(non-interface) value as
long as that value implements the  interface's methods.

An variable of interface type always has the same static type, and even though
at run time, the value stored in the interface variable may change type, that
value will always satisfy the interface.


The repesentation of an interface
----------------------------------
A variable of interfface type stores a pair: the concrete value assigned to the
variable, and that value's type descriptor. To be more precise, the value is
the underlying concrete data item that implements the inteface and the type
describes the full type of that item.

The static type of the interface determines what methods may be invoked with an
interface variable, even though the concrete value inside may have a larger set
of methods.

One important detail is that the pair inside an interface always has the form
(value, concrete type) and cannot have the form (value, innterface type).
Interfaces do not hold interface values.

The laws
========

1. Reflection goes from interface value to reflection object.
-------------------------------------------------------------

At the basic level, reflection is just a mechanism to examine the type and value
pair stored inside an interface variable.

TypeOf() returns the reflection type of the value in the interface{}

When we call reflect.TypeOf(x), x is first stored in an empty interface, which
is then passed as the argument; reflect.Typeof unpacks that empty interface to
recover the type information.

To keep the API simple, the getter and setter methods of Value operate on the
largest type that can hold the value.

The Kind of a value of a reflection describes the underlying type, not the
static type.

2. Reflectioin goes from reflection object to interface value.
--------------------------------------------------------------

Like physical reflection, refleciton in Go generats its own inverse.

Given a reflect.Value we can recover an interface value using the Interface
method; in effect the method packs the type and value information back into an
interface an interfface representation and the returns the result.

The Interface method is the inverse of the ValueOf function, except that its
result is always of static type interface{}

3. To modify a reflection object, the value must be settable.
-------------------------------------------------------------

Settability is a property of a reflection Value, and not all reflection Values
have it.

Settability is like addressability, but stricter. It is the property that a
reflection object can modify the actual storage that was used to create the
reflection object. Settabilit is determmined by whether the reflection object
holds the original item.

Reflection Values need the address of something in order to modify what they
represent.

*******************************************************************************/

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// 1st law
	{
		var x float64 = 3.4
		fmt.Printf("type:  %v, %#v \n", reflect.TypeOf(x), reflect.TypeOf(x))
		fmt.Printf("value: %v, %#v \n", reflect.ValueOf(x), reflect.ValueOf(x))

		v := reflect.ValueOf(x)
		fmt.Println("type: ", v.Type())
		fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
		fmt.Println("value:", v.Float())

		var y uint8 = 'x'
		v = reflect.ValueOf(y)
		fmt.Println("type: ", v.Type())
		fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8)
		y = uint8(v.Uint())

		// The Kind cannot discrimiate an int from a MyInt even though that
		// type can.
		type MyInt int
		var z MyInt = 7
		v = reflect.ValueOf(z)
		fmt.Println("type is: ", v.Type())
		fmt.Println("kind is:", v.Kind())
		PrintLine()
	}

	// 2nd law
	{
		var x float64 = 3.4
		v := reflect.ValueOf(x)
		y := v.Interface().(float64)
		fmt.Println(y)
		fmt.Printf("value is: %7.1e\n", v.Interface())
		PrintLine()
	}

	// 3rd law
	{
		//var x float64 = 3.4
		//v := reflect.ValueOf(x)
		//v.SetFloat(7.1)
		// panic: reflect: reflect.Value.SetFloat using unaddressable value

		var x float64 = 3.4
		v := reflect.ValueOf(x)
		fmt.Println("settability of v: ", v.CanSet())

		p := reflect.ValueOf(&x)
		fmt.Println("type of p:", p.Type())
		fmt.Println("settability of p:", p.CanSet())

		v = p.Elem()
		fmt.Println("settability of v:", v.CanSet())
		v.SetFloat(7.1)
		fmt.Println(v.Interface())
		fmt.Println(x)
		PrintLine()
	}

	// struct
	{
		type T struct {
			A int
			B string
		}
		t := T{23, "skidoo"}
		s := reflect.ValueOf(&t).Elem()
		typeOfT := s.Type()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
		s.Field(0).SetInt(77)
		s.Field(1).SetString("Sunset strip")
		fmt.Println("t is now: ", t)
	}
}

func PrintLine() {
	fmt.Println(strings.Repeat("-", 80))
}
