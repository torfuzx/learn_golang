/*
Various operation on slice.
*/

package main

import (
	"fmt"
)

func main() {
	{
		printLine()
		fmt.Println("Append Vector")

		a := []int{1, 2, 3, 4}
		b := []int{55, 66, 77}
		a = append(a, b...)
		fmt.Println(a)
	}

	{
		printLine()
		fmt.Println("Copy")

		// one way
		a := []int{1, 2, 3, 4}
		b1 := make([]int, len(a))
		copy(b1, a)
		fmt.Println(b1)

		// another way
		b2 := append([]int(nil), a...)
		fmt.Println(b2)
	}

	{
		printLine()
		fmt.Println("Cut")

		slice := []string{"s", "a", "v", "e", "k", "k", "k", "k", "k", "k", "m", "e"}
		slice = append(slice[:4], slice[10:]...)
		fmt.Printf("%v\n", slice)
	}

	{
		printLine()
		fmt.Println("Delete:")

		// method 1
		slice := []string{"good1", "good2", "good3", "bad1", "bad2", "good4", "good5"}
		slice2 := append(slice[:3], slice[5:]...)
		fmt.Println(slice2)

		// method 2
		// first move the elements after the remove elements foward until the
		// first element before the to-to-removed elements, and then append.
		slice3 := slice[:2+copy(slice[2:], slice[2+2:])]
		fmt.Println(slice3)

	}

	{
		//		printLine()
		//		println("Delete without preserving order:")
		//		slice := []string{"a1", "a2", "a3", "a4", "b1", "b2", "a5", "a6", "a7"}
		//		slice[4], slice = slice[len(slice)-1], slice[:len(slice)-1]
		//		fmt.Println(slice)
	}

	{
		printLine()
		fmt.Println("Expand:")
		slice := []int{1, 2, 3, 4, 5}
		fmt.Printf("slice len=%d, cap=%d, value=%v\n", len(slice), cap(slice), slice)
		slice2 := append(slice[:2], append(make([]int, 5), slice[2:]...)...)
		fmt.Printf("slice2 len=%d, cap=%d, value=%v\n", len(slice2), cap(slice2), slice2)
	}

	{
		printLine()
		println("Extend:")
		slice := []int{1, 2, 3, 4, 5}
		fmt.Printf("before extend: len=%d, cap=%d, value=%v\n", len(slice), cap(slice), slice)
		slice = append(slice, make([]int, 5)...)
		fmt.Printf("after extend: len=%d, cap=%d, value=%v\n", len(slice), cap(slice), slice)
	}

	{
		printLine()
		fmt.Println("Insert:")

		// remember the left inclusive and right exclusive rule is important
		slice := []int{1, 2, 3, 4, 7, 8, 9}
		slice = append(slice[:4], append([]int{5, 6}, slice[4:]...)...)
		fmt.Println(slice)

		// another way that avoids creating the new slice
		i := 4
		slice = []int{1, 2, 3, 4, 5, 7, 8, 9}
		slice = append(slice, 0)     // make a hole at the end
		copy(slice[i+1:], slice[i:]) // shift all the elements behind the position i by 1, and fill the hole
		slice[i] = 6                 // the i becomes a hole, and that's where we fill the new insert value
	}

	{
		// Note the i should the the index should the index of the first element
		// if new inserted elements
		printLine()
		fmt.Println("Insert vector:")
		slice1 := []int{1, 2, 3, 6, 7, 8, 9}
		slice2 := []int{4, 5}
		i := 3
		slice1 = append(slice1[:i], append(slice2, slice1[i:]...)...)
		fmt.Println(slice1)
	}

	{
		printLine()
		fmt.Println("Pop:")

		slice := []int{1, 2, 3, 4, 6, 7, 8}
		var x int // used to save the poped element

		// this will pop the last element from the slice
		x, slice = slice[len(slice)-1], slice[:len(slice)-1]
		fmt.Println("the poped element: ", x)
		fmt.Println("the slice after pop: ", slice)
	}

	{
		printLine()
		fmt.Println("Push:")
		slice := []int{1, 2, 3, 4, 5, 6, 7}
		x := 8
		slice = append(slice, x)
		fmt.Println(slice)
	}
}

func printLine() {
	fmt.Println("-------------------------------------------------------------")
}
