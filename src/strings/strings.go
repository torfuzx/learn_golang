package main

import (
	"fmt"
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main () {
	// LITERALS
	// --------
	{
		text1 := "\"what's that?\", he said"		// interpreted string literal
		text2 := `"what's that?", he said`			// raw strinf literal
		radicals := "√ \u221A \U0000221a"	// square root

		fmt.Println(text1)
		fmt.Println(text2)
		fmt.Println(radicals)
	}

	// COMPARE STRINGS
	// ---------------
	{
		book := "The Spirit Level" +
				" by Richard Wilkinson"
		book += " and Kate Pickett"
		fmt.Println(book)
		fmt.Println("Josey" < "José", "Josey" == "José")
	}

	// Convert a single character to a one-charater string
	{
		æs := ""
		for _, char := range []rune {
			'æ',
			0xE6,
			0346,
			230,
			'\xE6',
			'\u00E6',
		} {
			fmt.Printf("[0x%X '%c'] ", char, char)
			æs += string(char)
		}
		fmt.Println("");
		fmt.Println(æs)
	}

	{
		var buffer bytes.Buffer
		for _, piece := range []rune {
			'æ',
			0xE6,
			0346,
			230,
			'\xE6',
			'\u00E6',
		} {
			buffer.WriteString(string(piece))
		}
		fmt.Print(buffer.String(), "\n")
	}

	{
		phrase := "vått og tørt"
		fmt.Printf("string: \"%s\"\n", phrase)
		fmt.Println("index\trune\tchar\tbytes")
		for index, char := range phrase {
			fmt.Printf("%-2d\t%U\t'%c'\t%X\n", index, char, char, []byte(string(char)))
		}
	}

	{
		line := "røde og gule sløjfer"
		i := strings.Index(line, " ")
		firstWord := line[:i]
		j := strings.LastIndex(line, " ")
		lastWord := line[j+1:]
		fmt.Println(firstWord, lastWord)
	}

	{
		line := "rå tørt\u2028vær"
		i := strings.IndexFunc(line, unicode.IsSpace)
		firstWord := line[:i]
		j := strings.LastIndexFunc(line, unicode.IsSpace)
		_, size := utf8.DecodeRuneInString(line[j:])
		lastWord := line[j+size:]
		fmt.Println(firstWord, lastWord)
	}
}
