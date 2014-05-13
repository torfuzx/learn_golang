// hello.go
package main 
import {
	"fmt",
	"os",
	"strings"
}

func main () {
	who := "World!"
	if len(os.Args) > 1 {
		who = strings.Join(os.Arg[1:], " ")
	}
	fmt.println("Hello ", who)
}


