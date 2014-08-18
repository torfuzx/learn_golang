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
- The capacity is the length of the underlying array, minus the index in the
  array of the first element of the slice.

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

	var iBuffer [10]int
	iSlice := iBuffer[0:0]
	for i := 0; i < 20; i++ {
		iSlice = Extend(iSlice, i)
		fmt.Println(iSlice)
	}

	{
		printLine()
		slice := make([]int, 10, 15)
		fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

		// doubles the capacity but keeps the length the same
		newSlice := make([]int, len(slice), 2*cap(slice))
		for i := range slice {
			newSlice[i] = slice[i]
		}
		slice = newSlice
		fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	}

	{
		printLine()
		slice := make([]int, 10, 15)
		fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

		// doubles the capacity but keeps the length the same
		newSlice := make([]int, len(slice), 2*cap(slice))
		// copy the old data the right-hand argument to the left-hand argument
		copy(newSlice, slice)
		fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
		slice = newSlice
	}

	{
		printLine()
		slice := make([]int, 10, 20)
		for i := range slice {
			slice[i] = i
		}
		fmt.Println(slice)
		slice = Insert(slice, 5, 99)
		fmt.Println(slice)
	}

	{
		// test the robust version of extend
		printLine()
		slice := make([]int, 0, 5)
		for i := 0; i < 10; i++ {
			slice = Extend2(slice, 1)
			fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
			fmt.Println("address of 0th element: ", &slice[0])
		}
	}
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
	if cap(slice) == len(slice) {
		fmt.Println("slice is full, can't be extented.")
		return slice
	}

	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func printLine() {
	fmt.Println("-------------------------------------------------------------")
}

// Insert inserts the value into the slice at the specified index, which must be
// in range.
// The slice must have room for the new element.
func Insert(slice []int, index, value int) []int {
	// grow the slice by one element
	slice = slice[0 : len(slice)+1]
	// use copy to move the upper part of the slice out of the way and open a
	// hole
	copy(slice[index+1:], slice[index:])
	// store the new value
	slice[index] = value
	// return a result
	return slice
}

func Extend2(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		// slice is full, must grow
		// we double its size and add 1, so if the size is zero we still grow
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
