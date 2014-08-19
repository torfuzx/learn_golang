/*

RPC
---

Three restrictions:

- 1. The function must be public/exported (function name begin with upper case
  letter)
- 2. Have exactly two arguments, the first being a pointer to value data to be
  received by the function from the clieent, and the second a pointer to hold
  the returned answer to be sent back to the client
- 3. Have a return value of type error


*/

package main

import (
	"net/rpc"
	"errors"
	"net/http"
	"fmt"
)

type Args struct {
	A, B int
}

// Quotient and remainder
type Quotient struct {
	Quo, Rem int
}

// The following two functions will be registered with the RPC system, and
// to be called on the RPC server.
// Note they follow the 3 rules mentioned above.
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
