package main

import (
	"fmt"
	"unicode"
)

func main () {
	fmt.Println(IsHexDigit('8'), IsHexDigit('X'), IsHexDigit('x'), IsHexDigit('b'), IsHexDigit('B'))
}

func IsHexDigit(char rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, char)
}