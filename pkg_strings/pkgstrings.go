package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	{
		names := "Niccolò•Noël•Geoffrey•Amélie••Turlough•José"

		// Split
		fmt.Print("|")
		for _, name := range strings.Split(names, "•") {
			fmt.Printf("%s|", name)
		}
		fmt.Println()

		// SplitN
		fmt.Print("|")
		for _, name := range strings.SplitN(names, "•", 2) {
			fmt.Printf("%s|", name)
		}
		fmt.Println()

		fmt.Print("|")
		for _, name := range strings.SplitN(names, "•", 10) {
			fmt.Printf("%s|", name)
		}
		fmt.Println()

		// SplitAfter
		fmt.Print("|")
		for _, name := range strings.SplitAfter(names, "•") {
			fmt.Printf("%s|", name)
		}
		fmt.Println()

		// SplitAfterN
		fmt.Print("|")
		for _, name := range strings.SplitAfterN(names, "•", 2) {
			fmt.Printf("%s|", name)
		}
		fmt.Println()
	}

	{
		for _, record := range []string{
			"László Lajtha*1892*1963",
			"Édouard Lalo\t1823\t1892",
			"José Ángel Lamas|1775|1814",
		} {
			fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
				switch char {
				case '\t', '*', '|':
					return true
				}
				return false
			}))
		}
	}

	{
		names := " Antônio\tAndré\tFriedrich\t\t\tJean\t\tÉlisabeth\tIsabella \t"
		names = strings.Replace(names, "\t", " ", -1)
		fmt.Printf("|%s|\n", names)
		fmt.Printf("|%s|\n", SimpleSimplifyWhiteSpace(names))
		fmt.Printf("|%s|\n", SimplifyWhiteSpace(names))
	}

	{
		asciiOnly := func(char rune) rune {
			if char > 127 {
				return '?'
			}
			return char
		}
		removeNonAscii := func(char rune) rune {
			if char > 127 {
				return -1
			}
			return char
		}
		fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))
		fmt.Println(strings.Map(removeNonAscii, "Jérôme Österreich"))
	}
}

func SimpleSimplifyWhiteSpace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

func SimplifyWhiteSpace(s string) string {
	var buffer bytes.Buffer
	skip := true

	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			}
		} else {
			buffer.WriteRune(char)
			skip = false
		}
	}

	s = buffer.String()
	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}
