/*

Strings
--------

- String is in effect a read-only slice of bytes.
- String holds arbitrary bytes.
- Indexing a string yields its bytes, not characters.
- When we store a character value in a string, we store its byte-at-a-time
  representation
- Go source code is UTF-8, so the source code for the string literal is UTF-8
  text. If that string literal contains no escape sequences, which a raw string
  cannot, the constructed string will hold exactly the source text between the
  quotes. Thus by definition and by construction, the raw string will always
  contain a valid UTF-8 representation of its contents. Similarly, unless it
  contains UTF-8-breaking escapes, a regular string literal will also contai
  valid UTF-8
- Strings can contain arbitrary bytes, but when constructed from string literals
  those bytes are(almost always UTF=8)
- A string literal, absent byte-level escapes, always holds valid UTF8 sequences
- Those sequences represent Unicode code points, called runes
- No guarantee is made in Go that characters in strings are Normalized.

Rune
----
- Rune means the same as 'code point'
- Go define teh rune as an alias of type int32
- What you might think of as a character constant is called a rune constant
- A for-range loop, unlike for loop, decodes one UTF-8 encoded rune on each
  iteration. Each time aound the loop, the index of the loop is the starting
  of the current rune, measured in bytes, and the code point is its value.
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
		println("-------------------------------------------------------------")
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

		fmt.Println("Println:")
		fmt.Println(sample_slice)

		fmt.Println("Byte loop with %x:")
		for i := 0; i < len(sample_slice); i++ {
			fmt.Printf("%x ", sample_slice[i])
		}
		fmt.Printf("\n")

		// the %q will let the output be double-quoted string with Go syntax
		fmt.Println("Byte loop with %q:")
		for i := 0; i < len(sample_slice); i++ {
			fmt.Printf("%q ", sample_slice[i])
		}
		fmt.Printf("\n")

		fmt.Println("Print with %x:")
		fmt.Printf("%x\n", sample_slice)

		fmt.Println("Print with % x:")
		fmt.Printf("% x\n", sample_slice)

		// double quoted string
		fmt.Println("Print with %q:")
		fmt.Printf("%q\n", sample_slice)

		fmt.Println("Print with %+q:")
		fmt.Printf("%+q\n", sample_slice)

		// --------------------------------------------------------------------
		// Third set of examples
		// ---------------------------------------------------------------------
		println("-------------------------------------------------------------")
		const placeOfInterest = `⌘`

		fmt.Printf("plain string: ")
		fmt.Printf("%s", placeOfInterest)
		fmt.Printf("\n")

		fmt.Printf("quoted string: ")
		fmt.Printf("%+q", placeOfInterest)
		fmt.Printf("\n")

		fmt.Printf("hex bytes: ")
		for i := 0; i < len(placeOfInterest); i++ {
			fmt.Printf("%X ", placeOfInterest[i])
		}
		fmt.Printf("\n")

		// --------------------------------------------------------------------
		// Forth set of examples
		// ---------------------------------------------------------------------
		println("-------------------------------------------------------------")
		const name = "中华\u0000人民共和国"
		for index, runeValue := range name {
			fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
		}
	}

	// String literals
	// ---------------
	{
		println("-------------------------------------------------------------")
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
