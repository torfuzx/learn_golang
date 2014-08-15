package main

import (
	"fmt"
	"strings"
	"reflect"
)

type  BitFlag int
const (
	Active BitFlag = 1 << iota						// 1 << 0 == 1
	Send		// implictly 'BitFlag = 1 << iota'	// 1 << 1 == 2
	Receive		// implictly 'BitFlag = 1 << iota'	// 1 << 2 == 4
)
func (flag BitFlag) String() string {
	var flags []string

	if flag & Active == Active {
		flags = append(flags, "Active")
	}

	if flag & Send == Send {
		flags = append(flags, "Send")
	}

	if flag & Receive == Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 {// int(flag) is used to prevent infinite loop, very important
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return  "0()"
}

func main () {
	// VARIABLES & CONSTANTS
	// ---------------------
	const limit = 512
	const top uint16 = 1421
	start := -19
	end := int64(9876543210)
	var i int
	var debug = false
	checkResults := false
	stepSize := 1.5
	acronym := "FOSS"

	fmt.Println("type of 'limit':        ", reflect.TypeOf(limit))
	fmt.Println("type of 'to':           ", reflect.TypeOf(top))
	fmt.Println("type of 'start':        ", reflect.TypeOf(start))
	fmt.Println("type of 'end':          ", reflect.TypeOf(end))
	fmt.Println("type of 'i':            ", reflect.TypeOf(i))
	fmt.Println("type of 'debug':        ", reflect.TypeOf(debug))
	fmt.Println("type of 'checkResults': ", reflect.TypeOf(checkResults))
	fmt.Println("type of 'stepSize':     ", reflect.TypeOf(stepSize))
	fmt.Println("type of 'acronym':      ", reflect.TypeOf(acronym))

	// ENUERARTIONS
	// ------------
	// the following three groups has the same effect
	{
		{ // 1.
			const Cyan 		= 0
			const Magenta 	= 1
			const Yellow	= 2
			fmt.Println("First const declaration method - Cyan:%d Magenta:%d Yellow:%d", Cyan, Magenta, Yellow)
		}
		{ // 2.
			const (
				Cyan	= 0
				Magenta = 1
				Yellow	= 2
			)
			fmt.Println("Second const declaration method - Cyan:%d Magenta:%d Yellow:%d", Cyan, Magenta, Yellow)
		}
		{ // 3.
			const (
				Cyan = iota	// 0
				Magenta		// 1
				Yellow		// 2
			)
			fmt.Sprintln("Third declaration method - Cyan:%d Magenta:%d Yellow:%d", Cyan, Magenta, Yellow)
		}
	}

	flag := Active | Send
	fmt.Println(BitFlag(0), Active, Send, Receive, flag, flag|Receive)

	// BOOLEAN VALUES & EXPRESSIONS
	// ----------------------------



	// NUMERIC TYPES
	// -------------
	{
		const factor = 3
		i := 20000
		i *= factor
		j := int16(20)
		i += int(j)
		k := uint8(0)
		k = uint8(i)
	}

	// INTEGER TYPE
}




