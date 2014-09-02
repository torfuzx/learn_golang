/*******************************************************************************

Go interfaces (by Ian Lance Taylor)
[Source: http://www.airs.com/blog/archives/277]

The fact that you don't need to declare whether a type implements an interface
means that Go implements aa form of duck typing. This is not pure duck typing,
becase when possible the Go compiler will statically check whether the type
implementsf the interface. However, Go does have a purely dynamic aspect, in
that conversion is cehcked at runtime. If the conversion is invalid - if the
type of the value stored in the existing interface value doesn't satisfy the
interface to which it is being converted - the program will fail a runtime
error.

For every type which is converted to an interface type, gccgo will build a
type descriptor. Type descriptor are used by the type reflection support in the
reflect package, and they are also used by the internal runtime support. For
each type they define a (maybe empty) list of methods. For each method, five
piece of information aree stored:
1. a hash code for the type of the method (i.e: the parameters and result
  types)
2. the name of the method
3. for a method which is not exported, the name of the package in which it is
  defined
4. the type descriptor for the type of the method
5. a pointer to the function which implements the method

In gccgo, an interface value is really a struct with three fields;
1.a pointer to the type descriptor for the type of the value currently stored
  in the interface value
2.a pointer to a table of functions implementing the methods for the current
  value
3. a pointer to the current value itself

The table of functions is stored sorted by the name of the method, although
the name does not appear in the table. Calling a method on an interface means
calling a specific entry in this table, where the entry to call is known at
compile time. Thus the table of function is essentially the same as a C++
virtual function table, and calling a method on an interface requires the
same series of steps as calling a C++ virtual function: load the address of the
virtual table, load a specific entry in the table, and call it.

When a value is satically converted to an interface type, the gccgo compiler
will build the tables of the methods required for the value and that interface
type. This table is specific to the pair of types involved. gccgo will give this
table comdat linkage, so that it is only built once for each pair of types in
a program. Thus a static conversion to an interfface simply requires loading the
three fields of the interface struct with vlaues known at compile time.

A dynamic conversion from one interface type to another is more complex. Of the
three fields in the interface value, the type descriptor and the pointeer to the
real value can simply be copied to the new interface value. However. the  table
of methods must be built at runtime. This is done by looking at the list of
methods defined in the value's type descriptors and the list of methods defined
in the type descriptor for the interfface type itself. Both lists are sorted by
the name of the method. The runtime code merges the two sorted lists to produce
the method table. The merging is done using the name of the method and the type
has code. If the interface requires a method which the type doesn't provide, the
conversion fails.

Interface in Go are similiar to ideas in serveral other programming languages:
pure abstract virtual base classes in C++; typeclasses in Haskell; duck typing
in Python; etc. That said, i am not aware of other language which combimes
interfface values, static type checking, dynamic runtimee conversion, and no
requirement for explicit declaring that aa type satisfy an interface. The result
in Go is powerful, flexible, efficient and easy to  write.


*******************************************************************************/

package main

import (
	"fmt"
)

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}

func (p *S) Put(v int) {
	p.i = v
}

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	fmt.Println(p.Get())
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
	// panic: interface conversion: main.S is not main.I: missing method Get
}

func main() {
	var s S
	f(&s)

	h()

}
