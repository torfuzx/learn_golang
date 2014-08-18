/*

Strings
--------

- String is in effect a read-only slice of bytes.
- String holds arbitrary bytes.


*/

package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	// Based on the go blog: Strings, bytes, runes and characters in Go
	{
		// --------------------------------------------------------------------
		// First set of examples that work on strings
		// ---------------------------------------------------------------------
		const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
		fmt.Println("Println:")
		fmt.Println(sample)

		// index a string access individual bytes, not characters.
		fmt.Println("Byte loop:")
		for i := 0; i < len(sample); i++ {
			fmt.Printf("%x ", sample[i])
		}
		fmt.Printf("\n")

		// dumps out the sequence bytes of the string as hexidecimal digits,
		// two per bytes.
		fmt.Println("Print with %x:")
		fmt.Printf("%x\n", sample)

		// use space flag in the format
		fmt.Println("Print with % x:")
		fmt.Printf("% x\n", sample)

		// escape any non-printable byte sequence in a string so that the output
		// is not unambigous
		fmt.Println("Print with %q:")
		fmt.Printf("%q\n", sample)

		// escape not only non-printable sequences, but also any non-ASCII bytes
		// all while interpreting UTF-8. The result is that it expose the
		// Unicode value of properly formatted UTF-8 that represents non-ASCII
		// data in the string.
		fmt.Println("Print with %+q:")
		fmt.Printf("%+q\n", sample)

		// ---------------------------------------------------------------------
		// Second set of examples work on byte slices
		// Everything is the same as above, except that the string is coverted
		// to byte slice and every example is using slice of byte.
		// ---------------------------------------------------------------------
		println("-------------------------------------------------------------")
		sample_slice := []byte(sample)
		fmt.Println("Byte loop with %x:")
		for i := 0; i < len(sample_slice); i++ {
			fmt.Printf("%x ", sample_slice[i])
		}
		fmt.Printf("\n")

		fmt.Println("Byte loop with %q:")
		for i := 0; i < len(sample_slice); i++ {
			fmt.Printf("%q ", sample_slice[i])
		}
		fmt.Printf("\n")

		fmt.Println("Print with %x:")
		fmt.Printf("%x\n", sample_slice)

		fmt.Println("Print with % x:")
		fmt.Printf("% x\n", sample_slice)

		fmt.Println("Print with %q:")
		fmt.Printf("%q\n", sample_slice)

		fmt.Println("Print with %+q:")
		fmt.Printf("%+q\n", sample_slice)
		println("-------------------------------------------------------------")
	}

	// String literals
	// ---------------
	{
		text1 := "\"what's that?\", he said" // interpreted string literal
		text2 := `"what's that?", he said`   // raw strinf literal
		radicals := "√ \u221A \U0000221a"    // square root

		fmt.Println(text1)
		fmt.Println(text2)
		fmt.Println(radicals)
	}

	// Compare strings
	// ---------------
	{
		// concatenate strings
		book := "The Spirit Level" +
			" by Richard Wilkinson"
		book += " and Kate Pickett"
		fmt.Println(book)

		fmt.Println("Josey" < "José", "Josey" == "José")
	}

	// Convert a single character to a one-character string
	{
		æs := ""
		for _, char := range []rune{
			'æ',
			0xE6,
			0346,
			230,
			'\xE6',
			'\u00E6',
		} {
			// %X base 16, in lowercase.
			// %c the character represented by the coreesponding Unicode code point.
			fmt.Printf("[0x%X '%c'] ", char, char)
			æs += string(char)
		}
		fmt.Println("")
		fmt.Println(æs)
	}

	{
		var buffer bytes.Buffer
		for _, piece := range []rune{
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
