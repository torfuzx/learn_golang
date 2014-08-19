package main

import (
	"strconv"
	"fmt"
)

func main() {
	{
		for _, truth := range []string {"1", "t", "TRUE", "false", "F", "0", "5"} {
			if b, err := strconv.ParseBool(truth); err != nil {
				fmt.Printf("\n{%v}", err)
			} else {
				fmt.Print(b, " ")
			}
		}
		fmt.Println()
	}

	{
		x, err := strconv.ParseFloat("-99.7", 64)
		fmt.Printf("%8T %6v %v\n", x, x, err)

		y, err := strconv.ParseInt("71309", 10, 0)
		fmt.Printf("%8T %6v %v\n", y, y, err)

		z, err := strconv.Atoi("71309")
		fmt.Printf("%8T %6v %v\n", z, z, err)

		s := strconv.FormatBool(z > 100)
		fmt.Println(s)

		i, err := strconv.ParseInt("0xDEED", 0, 32)
		fmt.Println(i, err)

		j, err := strconv.ParseInt("0707", 0, 32)
		fmt.Println(j, err)

		k, err := strconv.ParseInt("10111010001", 2, 32)
		fmt.Println(k, err)
	}

	{
		i := 1679023
		fmt.Println(strconv.Itoa(i))
		fmt.Println(strconv.FormatInt(int64(i), 10))
		fmt.Println(strconv.FormatInt(int64(i), 2))
		fmt.Println(strconv.FormatInt(int64(i), 16))
	}

	{
		s := "Alle ønsker å være fri."
		quoted := strconv.Quote(s)
		fmt.Println(quoted)
		fmt.Println(strconv.Unquote(quoted))

	}

}