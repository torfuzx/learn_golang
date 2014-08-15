package main

import (
	"fmt"
	"math"
	"image/color"
)

type composer struct {
	name			string
	birthYear 		int
}

type rectangle struct {
	x0, y0, x1, y1 int
	fill color.RGBA
}

func main() {
	{
		y := 1.5
		fmt.Printf("y: %p -> %f\n", &y, y)

		y++
		fmt.Printf("y: %p -> %f\n", &y, y)

		z := math.Ceil(y)
		fmt.Printf("z: %p -> %f\n", &z, z)
		fmt.Println()
	}

	{
		x := 3
		y := 22
		fmt.Printf("x: %p -> %d; y: %p -> %d\n", &x, x, &y, y)

		px := &x
		fmt.Printf("x: %p -> %d; y: %p -> %d; px: %v; *px: %d\n", &x, x, &y, y, px, *px)
		x++
		fmt.Println("increnment x by 1")
		fmt.Printf("x: %p -> %d; y: %p -> %d; px: %v; *px: %d\n", &x, x, &y, y, px, *px)
		*px++
		fmt.Println("increnment the value pointed by px by 1")
		fmt.Printf("x: %p -> %d; y: %p -> %d; px: %v; *px: %d\n", &x, x, &y, y, px, *px)

		py := &y
		fmt.Printf("x: %p -> %d; y: %p -> %d; py: %v; *py: %d\n", &x, x, &y, y, py, *py)
		*py ++
		fmt.Printf("x: %p -> %d; y: %p -> %d; py: %v; *py: %d\n", &x, x, &y, y, py, *py)
		fmt.Println()
	}

	{
		z := 37
		pi := &z
		ppi := &pi
		fmt.Printf("|z  | %p | %12v |\n", &z, z)
		fmt.Printf("|pi | %p | %12v |\n", &pi, pi)
		fmt.Printf("|ppi| %p | %12v |\n", &ppi, ppi)

		**ppi++
		fmt.Printf("|z  | %p | %12v |\n", &z, z)
		fmt.Printf("|pi | %p | %12v |\n", &pi, pi)
		fmt.Printf("|ppi| %p | %12v |\n", &ppi, ppi)
		fmt.Println()
	}

	{
		i := 9
		j := 5
		product := 0

		swapAndProduct1(&i, &j, &product)
		fmt.Println(i, j, product)
	}

	{
		i := 9
		j := 5
		product := 0

		i, j, product = swapAndProduct2(i, j, product)
		fmt.Println(i, j, product)
	}

	{
		antónio := composer{"António Teixeira", 1707}
		fmt.Println(antónio)

		anges := new(composer)		// pointer to composer
		anges.name, anges.birthYear = "Agnes Zimmermann", 1845

		julia := &composer{}
		julia.name, julia.birthYear = "Julia Ward Howe", 1819

		augusta := &composer{"Augusta Holmès", 1847}

		fmt.Println(anges, julia, augusta)

		fmt.Println()
	}

	{
		grades := []int {87, 55, 43, 71, 60, 43, 32, 19, 63}
		inflate(grades, 3)
		fmt.Println(grades)
	}

	{
		rect := rectangle{4, 8, 20, 10, color.RGBA{0xFF, 0, 0, 0xFF}}
		fmt.Println(rect)
		resizeRect(&rect, 5, 5)
		fmt.Println(rect)
		fmt.Printf("%#v\n", rect)

	}
}

func swapAndProduct1(x, y, product *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
	*product = *x * *y
}

func swapAndProduct2(x, y, product int) (int, int, int) {
	if x > y {
		x, y = y, x
	}
	return x, y, x * y
}

func inflate(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}

func resizeRect(rect *rectangle, width, height int) {
	(*rect).x1 += width
	(*rect).y1 += height
}