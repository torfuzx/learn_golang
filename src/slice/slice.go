/*

array
-----

- Important building block in Go
- The size is part of its type
- Most common purpose in Go is to hold storage for a slice

slice
-----
- A data structure describing a contiguous section of an array stored separately
  from the slice variable itself.
- A slice is not an array. It describes a piece of an array.
- Under the cover, it's a struct value holding a pointer and a length. It's not
  a pointer to a struct.
- Slice header contains a length and a pointer to the an element of an array
- To have a function modify the slice header, you can pass a pointer to the
  slice header.
*/

package main

import (
	"bytes"
	"fmt"
)

var buffer [256]byte

type path []byte

func main() {
	// even though the slice header is passed by value, the header includes a
	// pointer to elements of an array, so both the original slice header and
	// the copy of the header passed to the function describes the array
	slice := buffer[10:20]

	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println("before:", slice)
	AddOneToEachElement(slice)
	fmt.Println("after :", slice)

	fmt.Println("before: ", slice)
	PtrSubtractOneFromLength(&slice)
	fmt.Println("after : ", slice)

	pathName := path("/usr/bin/tso") // conversion from string to path
	pathName.TruncateAtFinalSlash()
	fmt.Printf("%s\n", pathName)

	pathName = path("/home/kurt/projects")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)

}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

// a example to have a function modify the slice header(here, the len)
func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

// the method could operate on value because the value receivere still point to
// the same underlying array
func (p path) ToUpper() {
	for i, b := range p {
		if b >= 'a' && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

// demonstrate on slice capacity
// extents the slice of ints by one element
func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

