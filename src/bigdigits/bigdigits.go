package main 
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string {
	{
		"    0   ",
		"   0 0  ",
		"  0   0 ",
		"  0   0 ",
		"  0   0 ",
		"   0 0  ",
		"    0   ",
	},
	{
		"    1    ",
		"   11    ",
		"    1    ",
		"    1    ",
		"    1    ",
		"    1    ",
		"   111   ",		
	},
	{
		"   222   ",
		"  2   2  ",
		"     2   ",
		"    2    ",
		"   2     ",
		"  2      ",
		"  22222  ",		
	},
	{
		"   333   ",
		"  3   3  ",
		"      3  ",
		"    33   ",
		"      3  ",
		"  3   3  ",
		"   333   ",
	},
	{
		"     4   ",
		"    44   ",
		"   4 4   ",
		"  4  4   ",
		"  44444  ",
		"     4   ",
		"     4   ",
	},
	{
		"  55555  ",
		"  5      ",
		"  5      ",
		"   555   ",
		"      5  ",
		"  5   5  ",
		"   555   ",
	},
	{
		"  6     ",
		"  6     ",
		"  6     ",
		"  6666  ",
		"  6   6 ",
		"  6   6 ",
		"  6666  ",
	},
	{
		"  777777  ",
		"       7  ",
		"      7   ",
		"     7    ",
		"    7     ",
		"   7      ",
		"   7      ",
	},
	{
		"   888   ",
		"  8   8  ",
		"  8   8  ",
		"   888   ",
		"  8   8  ",
		"  8   8  ",
		"   888   ",
	},
	{
		"  9999  ",
		" 9   9  ",
		" 9   9  ",
		"  9999  ",
		"     9  ",
		"     9  ",
		"     9  ",
	},
}

func main () {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// the first argument, the number to be converted & displayed, the zero-th argument being the excutable's name itself
	stringOfDigits := os.Args[1]

	// iterate the numbers from user input using for...range 
	for row := range bigDigits[0] {
		line := ""
		// get the corresponding segment of its big digit format for the current line
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
 			if 0 <= digit && digit <= 9 {
 				line += bigDigits[digit][row] + " "
 			} else {
 				log.Fatal("invalid whole number")
 			}
		}
		fmt.Println(line)
	}
}


















