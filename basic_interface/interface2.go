/**

Go interfaces
-------------
@see: http://www.airs.com/blog/archives/277

Every type has an interface, which is a set of methods defined for that type.

The fact that you do not declare whether a type implements an interface means
that Go implments a form of duck typing. This is not pure duck typing, because
when possible the Go compiler will statically check whether the type implements
the interface. However, Go does have a purely dynamic aspect, in that you can
convert from one interfface type to another. In the general case, that
conversion is checked at runtime. If the conversion is invalid - if the type of
value stored in the exiting interface value doesn't satisfy the interface to
which it is being converted - the program will fail with a runtime error.

For every type which is converted to an interface type, gccgo will build a type
descriptor. Type descriptors are used by the type reflection support in the
reflect package, and they are also used by the internal runtime support.

For each type they define a (possibly empty) list of methods. For each method,
five pieces of information are stored:

1. a hash code for the type of the method (ie., the parameter and result types)
2. the name of the method
3. for a methodd which is not exported, the name of the package in which it is
  defined
4. the type descriptor for the type of the method
5. a pointer to the function which implements the method


In gccgo, an interfface value is really a struct with three fields
1. a pointer to the type descriptor for the type of the value currently stored
in the interface value
2. a pointer to a table of functions implementing the methods for the current
   value currently stored in the interface value
3. a pointer to the current valuee itself.


Thee table of function is stored by the name of the method, although the name
does not appear in the table. Calling a method on an interface means calling a
specific entry in this table, where the entry to call is known at compile time.
Thus the table of functions is essentially the sames as a C++ virtual functiion
table, and calling aa method on an interface requires the same series off steps
as calling a C++ virtual function: loading the addreds of  the virtual table,
load a specific entry in the table and call it.

When a value is statically converted to an interface type, the gccgo compiler
will build the table of  the methods required for that value and that interface
type. This table is specific to the pair of types involved. gccgo will gives the
table comdat linkage, so that it is only built once for each pairr of types in
a program. Thus a static conversion to a interface simply requires loading three
fields of the interface struct with value known at compile time.

A dynamic conversioin from one interfface to another is more complex. Pf thee
three fields in a interface value, the type descriptor and the pointer to the
real value can simply be copied to the new interface vlaue. However, the table
of methods must be built at runtime. This is done by lookingg at the list of
methods defined in the value's type descriptors and the list of  method in the
type descriptor for the interfface itself. Both lists are sorted by the name of
the method. The runtime code merges the two sorted lists to produce the method
table. If the interface requires a method which the types doesn't provide, the
conversion fails.

Interface in Go are similiar to ideas in serverl other languages: pure abstract
virtual base classes in C++; typeclasses in Haskell; duck typing in Pyhon; etc.
That said, static type checking , dynamic runtime conversion, and no requirement
for explictly declaring that type satifies an interface. The result in Go is
powerful, flexible and easy to write.

*/

package main

import (
	"fmt"
)

type I interface {
	Get() int
	Put(int)
}

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}

func (p *S) Put(v int) {
	p.i = v
}

// Interface value, p holds a value of interface type
func f(p I) {
	p.Put(1)
	fmt.Println(p.Get())
}

func g(i interface{}) int {
	return i.(I).Get()
}

func h() {
	var s S
	fmt.Println(g(&s))
	fmt.Println(g(s))
}

func main() {
	var s S
	f(&s)
}
