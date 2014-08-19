/**

RPC Client
----------


The RPC client needs to set up a HTTP connection to the RPC server. It needs to
prepare a structure with the values to be sent, and the address of a variable
to store the results in. The it make a Call with arguments:

- The name of the remote function to execute
- The value to be sent
- The address of a variable to store the result in

*/

package main

import (
	"fmt"
	"os"
	"net/rpc"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddr := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddr+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d / %d = %d, remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
