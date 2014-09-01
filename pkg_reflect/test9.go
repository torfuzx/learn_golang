/*******************************************************************************

This is about interface's internal

Source: http://research.swtch.com/interfaces

Interface - static, checked at compile time, dynamic when asked for.

Languages with methods typically fall into one tof two camps:

- prepare tables for all the method statically
- do a method lookup at each call

Go sits halfway between the two: it has method tables but computes them at run
time.

Interface values are represented as a two-word pair giving a pointer to
information about the type stored in the interface and a pointer to the
associated data.

- The first word in the interface value points at what I call an interface tabble
or itable. The itable begins with some metadata about the types involded nad then
begins with some metadata about hte tyoes involved and then becomes a list of 
functioinn pointers.. NOte that the itable correspondes to the interfacee type,
not the dynamic type.

- The second word points at the actual data. 
  Value stored inn interffacces might be arbitrarily large, but only one word is dedicated to
  holding the vlaue n the interfface structure, so the assignment allocates
  a chunk of memory on the heap and records the pointer in the one-word slot.
  
  To check  whether an interface value holds a particular type, the go compiler
  generates code equivalent to the C expression s.tab->type to obtain the type 
  pointeer and check it against the desired type. If  the types match, the value
  can be copied by dereferencing s.data.
  
  				
  					s := Stringer(200)
 				tab	┌───────┐
  					│──200──│
  				data└───────┘ 					
  					Stringer
  									itable(Stringer, Binary)
  ┌───────┐							┌───────┐	
  │──200──│						type│───────│ type(Binary)
  └───────┘							│		│
  									│		│
  Binary							│───────│
  							fun[0]	└───────┘(*Binary).String
  									itable
  									(in C)Itab
  									

*******************************************************************************/

package main

import (
	"fmt"
)

type Binary uint64

type Stringer interface {
	String() string
}

func (i Binary) String() string {
	return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
}
