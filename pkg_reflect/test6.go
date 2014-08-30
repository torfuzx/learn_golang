package main

import (
	"fmt"
	"reflect"
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
	Greetings string
	Message   string
	Pi        float64
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

func main() {
	// Imagine we have no influence on the value returned by create
	created := create()

	{
		fmt.Println("Translating aaa struct")
		original := created.(B)
		translated := translate(original)
		fmt.Println("original: ", original, "->", original.Ptr)
		fmt.Println("translated: ", translated, "->", translated.(B).Ptr)
	}
	fmt.Println()
	{
		fmt.Println("Translating a struct wrapped in an interface")
		original := created
		translated := translate(original)
		fmt.Println("original: ", (*original), "->", original.Ptr)
		fmt.Println("translated: ", (*translated.(*I)), "->", (*translated.(*I)).(B).Ptr)
	}
	fmt.Println()
	{
		fmt.Println("Translating a pointer to a struct wrapped in a interface")
		original := &created
		translated := translate(original)
		fmt.Println("original: ", (*original), "->", (*original).(B).Ptr)
		fmt.Println("translated: ", (*translated).(*I), "->", (*translated.(*I)).(B).Ptr)
	}
	fmt.Println()
	{
		fmt.Println("Translating a struct containing a pointer to a struct wrapped in an interface")
		type D struct {
			Payload *I
		}
		original := D{
			Payload: &created,
		}
		translated := translate(original)
		fmt.Println("original: ", original, "->", (*original.Payload), "->", (*original.Payload).(B).Ptr)
		fmt.Println("translated:", translated, "->", (*translated.(D).Payload), (*(translated.(D).Payload)).(B).Ptr)
	}

}

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
				string: "deep",
			},
		},
		Slice: []string{
			"and once more",
		},
		Answer: 42,
	}
}

func translate(obj interface{}) interface{} {
	// wrap the original in a reflect.value
	original := reflect.ValueOf(obj)

	cpy := reflect.New(original.Type()).Elem()
	translateRecursive(cpy, original)

	// remove the  reflection wrapper
	return cpy.Interface()
}

func translateRecusive(cpy, original reflect.Value) {
	switch original.Kind() {
	// the first case handle nested structures and translate them recusively
	case reflect.Interface:
		// get rid of the wrapping interface
		originalValue := original.Elem()
		// now gives us a pointer, but we want the value it points to
		cpyValue := reflect.New(originalValue.Type()).Elem()
		translateRecusive(cpyValue, originalValue)
		cpy.Set(cpyValue)
	case reflect.Ptr:
		// to get the actual type of the original we have to call Elem()
		cpy.Set(reflect.New(original.Elem().Type()))
		// unwrap the pointers so we don't end up in a infinite recusion
		translateRecusive(cpy.Elem(), original.Elem())
	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			translateRecusive(cpy.Field(i), original.Field(i))
		}
	case reflect.Slice:
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			translateRecusive(cpy.Index(i), original.Index(i))
		}
	case reflect.Map:
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			// New gives us a pointer, but again we want the value
			cpyValue := reflect.New(originalValue.Type()).Elem()
			translateRecusive(cpyValue, originalValue)
			cpy.SetMapIndex(key, cpyValue)
		}
	// the last two cases finish the recusion
	case reflect.Reflect.String:
		translatedString := dict[original.Interface().(string)]
		cpy.SetString(translatedString)
	default:
		cpy.Set(original)
	}
}
