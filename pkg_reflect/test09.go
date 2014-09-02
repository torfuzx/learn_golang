/*******************************************************************************

This is about interface's internal

Source: http://research.swtch.com/interfaces

Interface - static, checked at compile time, dynamic when asked for.

Languages with methods typically fall into one of two categories:

- prepare tables for all the method statically(as in C++  and Java)
- do a method lookup at each call and add fancy caching to make that call
  efficient

Go sits halfway between the two: it has method tables but computes them at run
time.

Interface values
----------------

Interface values are represented as a two-word pair giving a pointer to
information about the type stored in the interface and a pointer to the
associated data.

- The first word in the interface value points at what I call an interface table
  or itable. The itable begins with some metadata about the types involved, then
  becomes a list of functioinn pointers. Note that the itable correspondes to
  the interface type, not the dynamic type.

- The second word points at the actual data.
  Value stored in interfaces might be arbitrarily large, but only one word is
  dedicated to holding the value in the interface structure, so the assignment
  allocates a chunk of memory on the heap and records the pointer in the
  one-word slot.

  To check  whether an interface value holds a particular type, the go compiler
  generates code equivalent to the C expression s.tab->type to obtain the type
  pointeer and check it against the desired type. If  the types match, the value
  can be copied by dereferencing s.data.

  To call s.String(), the Go compiler generates code that does the equivalent
  of the C expression s->tab->fun[0](s.data): it calls the appropriate function
  pointer from the itable, passing the interface value's data word as the
  function's first(in this example only) argument. Note that the function in
  the itable is being passed the 32-bit pointer from the second word of the
  interface value, not the 64-bit value it points to. In general, the interface
  call site doesn't know the meaning of this word nor how much data it points to
  . Instead, the interface code arranges that the function pointers in the
  itable expect the 32-bit representation stored in the interface values. Thus
  the functioin pointer in this example is (*Binary).String() not
  Binary.String()


  ------------------------------------------------------------------------------
                      b := Binary(200)
                       ┌───────┐
                       │──200──│
                       └───────┘
                        Binary
  ------------------------------------------------------------------------------

                  s := Stringer(b)
                tab ┌───────┐
                    │──200──│
               data └───────┘
                     Stringer
                                 itable(Stringer, Binary)
  ┌───────┐                         ┌───────┐
  │──200──│                     type│───────│ type(Binary)
  └───────┘                         │       │
                                    │       │
  Binary                            │───────│
                           fun[0]   └───────┘(*Binary).String
                                      itable
                                     (in C)Itab

  - The itable for Stringer holding type Binary lists the methods used to
    satisfy Stringer, which is just String(). Binary's other methods(Get)
    make no appearance in the itable.

  - The second word points to the actual data, in this case a copy of b.
    The assignment var s Stringer = b makes a copy of b rather than point at
    b for if later b changes, s are supposed to have the original value, not the
    new one.
  ------------------------------------------------------------------------------

	(the upper word points to the left, the lower word points to the right)


Computing the Itable
--------------------

Go's dynamic type conversion mean that it isn't reasonable for the compiler
or linker to precompute all poosible itables: there are too many(interface
type, concrete type)pairs, and most won't be needed. Instead the compiler
generates a type description structure for each concrete type like Binary or int
or func(map[int]string). Among other metadata, the type description structure
contains a list of the methods implemented by the type. Similarly, the compiler
generates a (different) type description structure for each interface type like
Stringer; it too contains a method list. The interface runtime computes the
itable by looking forr each method listed in the interface type's method table
in the concrete type's method table. The runtime caches the itable after
generating it, so that this correpindece need only be computed once.

In our simpile example, the method table for Stringer has one method, while
the table for Binary has two methods. In gerneral, there might be ni methods
for interface type and nt methods for the concrete type. The obvious search
to find the mapping from interface methods to concrete methods would take
O(ni x nt) time, but we can do better. By sorting the two method tables and
walking them simultaneously, we can build the mapping in O(ni + nt) time
instead.

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
