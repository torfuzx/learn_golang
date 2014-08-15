package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"math"
)

type polar struct {
	radius float64
	angle  float64
}

func main() {
	// boolean
	{
		fmt.Printf("%t %t\n", true, false)
		fmt.Printf("%d %d\n", IntForBool(true), IntForBool(false))
	}

	// integer
	{
		// binary
		fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37, 37, 37, 37, 37)

		// octal
		fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%#+08o|\n", 41, 41, 41, 41, 41)

		// hexidecimal
		i := 3931
		fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i)

		// decimal
		i = 569
		fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i, i, i, Pad(i, 6, '*'))
	}

	// strings
	{
		fmt.Printf("%d %#04d %U '%c'\n", 0x3A6, 934, '\u03A6', '\U000003A6')
	}

	// float
	{
		for _, x := range []float64 {-.258, 7194.84, -60897162.0218, 1.500089e-8} {
			fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ','))
		}
	}

	// string and slice
	{
		slogan := "End Óréttlæti♥"
		fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)

		chars := []rune(slogan)
		fmt.Printf("%x\n%#x\n%#X", chars, chars, chars)

		bytes := []byte(slogan)
		fmt.Printf("%s\n%x\n%X\n% X\n%v\n", bytes, bytes, bytes, bytes, bytes)

		s := "Dare to be naive"
		fmt.Printf("|%22s|%-22s|%10s|\n", s, s, s)

		i := strings.Index(s, "n")
		fmt.Printf("|%.10s|%.*s|%-22.10s|%s|\n", s, i, s, s, s)
	}

	// formatting for debugging
	{
		p := polar{-83.40, 71.60}
		fmt.Printf("|%T|%v|%#v|\n", p, p, p)
		fmt.Printf("|%T|%v|%t|\n", false, false, false)
		fmt.Printf("|%T|%v|%d|\n", 7607, 7607, 7607)
		fmt.Printf("|%T|%v|%f|\n", math.E, math.E, math.E)
		fmt.Printf("|%T|%v|%f|\n", 5+7i, 5+7i, 5+7i)

		s := "Relativity"
		fmt.Printf("|%T|\"%v\"|\"%s\"|\"%q\"|\n", s, s, s, s)
	}

	{
		s := "Alias↔Synonym"
		chars := []rune(s)
		bytes := []byte(s)
		fmt.Printf("%T: %v\n%T: %v\n", chars, chars, bytes, bytes)
	}

	{
		i := 5
		f := -48.3124
		s := "Tomás Bretón"
		fmt.Printf("| %p -> %d | %p -> %f | %#p -> %s |\n", &i, i, &f, f, &s, s)
	}

	{
		fmt.Println([]float64 {math.E, math.Pi, math.Phi})
		fmt.Printf("%v\n", []float64 {math.E, math.Pi, math.Phi})
		fmt.Printf("%#v\n", []float64 {math.E, math.Pi, math.Phi})
		fmt.Printf("%.5f\n", []float64 {math.E, math.Pi, math.Phi})

		fmt.Printf("%q\n", []string{"Software patents", "kill", "innovation"})
		fmt.Printf("%v\n", []string{"Software patents", "kill", "innovation"})
		fmt.Printf("%#v\n", []string{"Software patents", "kill", "innovation"})
		fmt.Printf("%17s\n", []string{"Software patents", "kill", "innovation"})

		fmt.Printf("%v\n", map[int]string{1:"A", 2:"B", 3:"C", 4:"D"})
		fmt.Printf("%#v\n", map[int]string{1:"A", 2:"B", 3:"C", 4:"D"})

		fmt.Printf("%v\n", map[int]int{1:1, 2:2, 3:4, 4:8})
		fmt.Printf("%#v\n", map[int]int{1:1, 2:2, 3:4, 4:8})
		fmt.Printf("%04b\n", map[int]int{1:1, 2:2, 3:4, 4:8})
	}

}

func IntForBool (b bool) int {
	if b {
		return 1
	}
	return 0
}

func Pad(number, width int, pad rune) string {
	s := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(s)

	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}
	return s
}


func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	// return the integer and fractional-point number that sums to amount
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:] // strip ±
	fraction := ""

	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents) // force unify format, apply the lengh rule which is speicified by the decimal param
		fraction = fraction[2:] // strip the leading ±0
	}

	// apply the thousand separator
	sep := string(separator) // convert rune/character to string
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]		// insert the separator
		fmt.Println("whole: ", whole)
	}

	// put back the minus sign if the number is negative
	if amount < 0.0 {
		whole = "-" + whole
	}

	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)

	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}