/*******************************************************************************

Adapted from:
http://www.snip2code.com/Snippet/47387/Golang-reflection--traversing-arbitrary-

*******************************************************************************/
package main

import (
	"fmt"
	"reflect"
	"strings"
)

var dict = map[string]string{
	"Hello!":                 "你好!",
	"What's up?":             "最近过得怎么样?",
	"translate this":         "把这个给我翻译了",
	"point here":             "指着这里",
	"translate this as well": "把这个也给我翻译了",
	"and once more":          "再给也来一遍",
	"deep":                   "要有深度",
}

type I interface{}

type A struct {
	Greeting string
	Message  string
	Pi       float64
}

type B struct {
	Struct    A
	Ptr       *A
	Answer    int
	Map       map[string]string
	StructMap map[string]C
	Slice     []string
}

type C struct {
	String string
}

type D struct {
	Payload *I
}

// Be noted that it returns an interface.
func create() I {
	return B{
		Struct: A{
			Greeting: "Hello",
			Message:  "translate this",
			Pi:       3.14,
		},
		Ptr: &A{
			Greeting: "What's up?",
			Message:  "point here",
			Pi:       3.14,
		},
		Map: map[string]string{
			"Test": "translate this",
		},
		StructMap: map[string]C{
			"C": C{
				String: "deep",
			},
		},
		Slice: []string{
			"and once more",
		},
		Answer: 42,
	}
}

func main() {
	// Imagine we have no influence on the value returned by create(), don't
	// what type it will return.
	created := create()
	printLine()

	{
		fmt.Println("# Translating a struct:")
		original := created.(B)
		translated := translate(original)
		printOutcome(original, original.Ptr, translated, translated.(B).Ptr)
		printLine()
	}
	{
		fmt.Println("# Translating a struct wrapped in an interface:")
		original := created
		translated := translate(original)
		printOutcome(original, original.(B).Ptr, translated, translated.(B).Ptr)
		printLine()
	}
	{
		fmt.Println("# Translating a pointer to a struct wrapped in a interface:")
		original := &created
		translated := translate(original)
		printOutcome((*original), (*original).(B).Ptr, (*translated.(*I)), (*translated.(*I)).(B).Ptr)
		printLine()
	}
	{
		fmt.Println("# Translating a struct containing a pointer to a struct wrapped in an interface: ")
		original := D{Payload: &created}
		translated := translate(original)
		printOutcome2(original, (*original.Payload), (*original.Payload).(B).Ptr,
			translated, (*translated.(D).Payload), (*(translated.(D).Payload)).(B).Ptr)
	}
}

func translate(obj interface{}) interface{} {
	// wrap the original in a reflect.value
	original := reflect.ValueOf(obj)

	// New()  - Returns a pointer to a new zeroed value for the specified value.
	// Elem() - Returns the value interface v contains or that poiner v points to.
	//          It panics if v's kind is not Interface or Ptr.
	cpy := reflect.New(original.Type()).Elem()
	translateRecursive(cpy, original)

	// remove the reflection wrapper
	// Interface() - Returns v's current value as an interface{}. It's
	//               equivalent to:
	//               var i interface{} = (v's underlying value)
	return cpy.Interface()
}

func translateRecursive(cpy, original reflect.Value) {
	switch original.Kind() {
	// the first case handle nested structures and translate them recusively
	case reflect.Interface:
		// get rid of the wrapping interface
		originalValue := original.Elem()
		// now gives us a pointer, but we want the value it points to
		cpyValue := reflect.New(originalValue.Type()).Elem()
		translateRecursive(cpyValue, originalValue)
		cpy.Set(cpyValue)

	case reflect.Ptr:
		// to get the actual type of the original we have to call Elem()
		cpy.Set(reflect.New(original.Elem().Type()))
		// unwrap the pointers so we don't end up in a infinite recusion
		translateRecursive(cpy.Elem(), original.Elem())

	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			translateRecursive(cpy.Field(i), original.Field(i))
		}

	case reflect.Slice:
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			translateRecursive(cpy.Index(i), original.Index(i))
		}

	case reflect.Map:
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			// New gives us a pointer, but again we want the value
			cpyValue := reflect.New(originalValue.Type()).Elem()
			translateRecursive(cpyValue, originalValue)
			cpy.SetMapIndex(key, cpyValue)
		}

	// the last two cases finish the recusion
	case reflect.String:
		translatedString := dict[original.Interface().(string)]
		cpy.SetString(translatedString)

	default:
		cpy.Set(original)
	}
}

func printLine() {
	fmt.Println(strings.Repeat("-", 80))
}

func printOutcome(a, b, c, d interface{}) {
	fmt.Printf("## original:\n-> %#v\n-> %#v\n", a, b)
	fmt.Printf("## translated:\n-> %#v\n-> %#v\n", c, d)
}

func printOutcome2(a, b, c, d, e, f interface{}) {
	fmt.Printf("## original:\n-> %#v\n-> %#v\n\n-> %#v\n", a, b, c)
	fmt.Printf("## translated:\n-> %#v\n-> %#v\n\n-> %#v\n", d, e, f)
}
