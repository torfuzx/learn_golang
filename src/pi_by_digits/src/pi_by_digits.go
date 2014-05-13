package main

import (

)

func main () {
	places := handleCommandLine(1000)
	scaledPi := fmt.Sprint(π(places))
	fmt.Printf("3.%s\n", scaledPi[1:])
}
