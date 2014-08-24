package main

import (
	"fmt"
)

type Product struct {
	name  string
	price float64
}

func main() {
	{
		// 1. declared as '[length]Type'
		var buffer [20]byte

		var grid1 [3][3]int
		grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2

		// 2. declared as '[N]Type{value1, ...., valueN}'
		grid2 := [3][3]int{{4, 3}, {8, 6, 2}}

		// 3. declared as '[...]Type{value1, ..., valueN}'
		cities := [...]string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
		cities[len(cities)-1] = "Karachi"

		fmt.Println("Type len contents")
		fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
		fmt.Printf("%-8T %2d %v\n", grid1, len(grid1), grid1)
		fmt.Printf("%-8T %2d %v\n", grid2, len(grid2), grid2)
		fmt.Printf("%-8T %2d %q\n", cities, len(cities), cities)
		fmt.Println()
	}

	{
		s := []string{"A", "B", "C", "D", "E", "F", "G"}
		t := s[:5] // left inclusive, right exclusive
		u := s[3 : len(s)-1]
		fmt.Println(s, t, u)
		u[1] = "X"
		fmt.Println(s, t, u)

		fmt.Println()
	}

	{
		s := new([7]string)[:]
		s[0], s[1], s[2], s[3], s[4], s[5], s[6] = "A", "B", "C", "D", "E", "F", "G"

		t := s[:5]
		u := s[3 : len(s)-1]

		fmt.Printf("s: len(%d) cap(%d) -> %#v\n", len(s), cap(s), s)
		fmt.Printf("t: len(%d) cap(%d) -> %#v\n", len(t), cap(t), t)
		fmt.Printf("u: len(%d) cap(%d) -> %#v\n", len(u), cap(u), u)

		fmt.Println()
	}

	{
		// syntax 1: make([]Type, length, capacity)
		buffer := make([]byte, 20, 60)

		// syntax 2: make([]Type, length)
		grid1 := make([][]int, 3)
		for i := range grid1 {
			grid1[i] = make([]int, 3)
		}

		grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
		grid2 := [][]int{{4, 3, 0}, {8, 6, 2}, {0, 0, 0}}
		grid3 := [][]int{{9, 7}, {8}, {4, 2, 6}}

		cities := []string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
		cities[len(cities)-1] = "Karachi"

		fmt.Println("Type Len Cap Contents")
		fmt.Printf("%-8T %2d %3d %#v\n", buffer, len(buffer), cap(buffer), buffer)
		fmt.Printf("%-8T %2d %3d %#v\n", cities, len(cities), cap(cities), cities)
		fmt.Printf("%-8T %2d %3d %#v\n", grid1, len(grid1), cap(grid1), grid1)
		fmt.Printf("%-8T %2d %3d %#v\n", grid2, len(grid2), cap(grid2), grid2)
		fmt.Printf("%-8T %2d %3d %#v\n", grid3, len(grid3), cap(grid3), grid3)

		fmt.Println()
	}

	{
		s := []string{"A", "B", "C", "D", "E", "F", "G"}
		t := s[2:6]
		fmt.Println(t, s, "=", s[:4], "+", s[4:])
		fmt.Printf("s len(%d) cap(%d)\n", len(s), cap(s))
		fmt.Printf("t len(%d) cap(%d)\n", len(t), cap(t))

		s[3] = "x"
		t[len(t)-1] = "y"
		fmt.Println(t, s, "=", s[:4], "+", s[4:])
		fmt.Printf("s len(%d) cap(%d)\n", len(s), cap(s))
		fmt.Printf("t len(%d) cap(%d)\n", len(t), cap(t))
	}

	{
		amounts := []float64{237.81, 261.87, 273.93, 279.99, 281.07, 303.17,
			231.47, 227.33, 209.23, 197.09}
		sum := 0.0
		for _, amount := range amounts {
			sum += amount
		}
		fmt.Printf("Σ %.1f → %.1f \n", amounts, sum)

		sum = 0.0
		for i := range amounts {
			amounts[i] *= 1.05
			sum += amounts[i]
		}
		fmt.Printf("Σ %.1f → %.1f \n", amounts, sum)
	}

	{
		products := []*Product{
			{"Spanner", 3.99},
			{"Wrench", 2.49},
			{"Screwdriver", 1.99},
		}
		fmt.Println(products)
		for _, product := range products {
			product.price += 0.50
		}
		fmt.Println(products)
		fmt.Println()
	}

	// modify slice
	{
		s := []string{"A", "B", "C", "D", "E", "F", "G"}
		t := []string{"K", "L", "M", "N"}
		u := []string{"m", "n", "o", "p", "q", "r"}

		s = append(s, "h", "i", "j") // add single value
		s = append(s, t...)          // add all the values in a slice
		s = append(s, u[2:5]...)     // add a sub-slice

		b := []byte{'U', 'V'}
		letters := "WXY"
		b = append(b, letters...) // add a character byte in to a slice of bytes

		fmt.Printf("%q\n%s\n", s, b)
		fmt.Println()
	}

	{
		s := []string{"M", "N", "O", "P", "Q", "R"}
		x := InsertStringSliceCopy(s, []string{"a", "b", "c"}, 0) // at the front
		y := InsertStringSliceCopy(s, []string{"x", "y"}, 3)      // in the middle
		z := InsertStringSliceCopy(s, []string{"z"}, len(s))      // at the end
		fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)
		fmt.Println()

		x = InsertStringSlice(s, []string{"a", "b", "c"}, 0) // at the front
		y = InsertStringSlice(s, []string{"x", "y"}, 3)      // in the middle
		z = InsertStringSlice(s, []string{"z"}, len(s))      // at the end
		fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)
		fmt.Println()
	}

	{
		s := []string{"A", "B", "C", "D", "E", "F", "G"}
		s = s[2:]
		fmt.Println(s)
		fmt.Println()

		s2 := []string{"A", "B", "C", "D", "E", "F", "G"}
		s2 = s2[:4]
		fmt.Println(s2)
		fmt.Println()

		s3 := []string{"A", "B", "C", "D", "E", "F", "G"}
		s3 = append(s3[:1], s3[5:]...)
		fmt.Println(s3)
		fmt.Println()
	}

	{
		s := []string{"A", "B", "C", "D", "E", "F", "G"}
		x := RemoveStringSliceCopy(s, 0, 2)
		y := RemoveStringSliceCopy(s, 1, 5)
		z := RemoveStringSliceCopy(s, 4, len(s))
		fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)
		fmt.Printf("%v\n", s)
		fmt.Println()
	}

	{
		s := []string{"A", "B", "C", "D", "E", "F", "G"}
		x := RemoveStringSlice(s, 0, 2)
		y := RemoveStringSlice(s, 1, 5)
		z := RemoveStringSlice(s, 4, len(s))
		fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)
		fmt.Printf("%v\n", s)
		fmt.Println()
	}

	{

	}
}

func (product Product) String() string {
	return fmt.Sprintf("%s (%.2f)", product.name, product.price)
}

func InsertStringSliceCopy(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}

func RemoveStringSliceCopy(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

func RemoveStringSlice(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}
