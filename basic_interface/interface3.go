/**

Go Data Structures: Interfaces
by Russ Cox, December 1, 2009
@see: 		
------------------------------------------

Languages with methods typically fall into one of two camps: prepare tables for
all the method calls statically, or do a method lookup at each call and add 
fancy cachingg to make that call efficient. Go sits halfway between the two: 
it has method table but computes them at runtime. 

Interface value are represented as a two-word pair giving a pointer  to 
information about the type stored in the interface and a pointer to the 
associated data. 

*/

package main

import (
	"fmt"
)

func main() {

}
