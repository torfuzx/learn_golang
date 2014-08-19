package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// bigDigits is is two dimentional slice, its type is: [][]string,
// or in simple words: a slice of (slice of type string)
var bigDigits = [][]string{
	{
		"  0  ",
		" 0 0 ",
		"0   0",
		"0   0",
		"0   0",
		" 0 0 ",
		"  0  ",
	},
	{
		"  1  ",
		" 11  ",
		"  1  ",
		"  1  ",
		"  1  ",
		"  1  ",
		" 111 ",
	},
	{
		" 222 ",
		"2   2",
		"   2 ",
		"  2  ",
		" 2   ",
		"2    ",
		"22222",
	},
	{
		" 333 ",
		"3   3",
		"    3",
		"  33 ",
		"    3",
		"3   3",
		" 333 ",
	},
	{
		"   4 ",
		"  44 ",
		" 4 4 ",
		"4  4 ",
		"44444",
		"   4 ",
		"   4 ",
	},
	{
		"55555",
		"5    ",
		"5    ",
		" 555 ",
		"    5",
		"5   5",
		" 555 ",
	},
	{
		" 666 ",
		"6    ",
		"6    ",
		"6666 ",
		"6   6",
		"6   6",
		"6666 ",
	},
	{
		"777777",
		"     7",
		"    7 ",
		"   7  ",
		"  7   ",
		" 7    ",
		" 7    ",
	},
	{
		" 888 ",
		"8   8",
		"8   8",
		" 888 ",
		"8   8",
		"8   8",
		" 888 ",
	},
	{
		" 9999",
		"9   9",
		"9   9",
		" 9999",
		"    9",
		"    9",
		" 999 ",
	},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// fmt.Printf("type of variable bigDigits: %s\n", reflect.TypeOf(bigDigits))

	// the first argument(index value 0) being the excutable's name itself
	// the second argument, the number to be converted & displayed
	stringOfDigits := os.Args[1] // type: string

	rowNum := len(bigDigits[0])
	for rowIdx := 0; rowIdx < rowNum; rowIdx++ {
		line := ""
		for _, digit := range stringOfDigits {
			digit -= '0'

			if digit >= 0 && digit <= 9 {
				line += bigDigits[digit][rowIdx] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}
