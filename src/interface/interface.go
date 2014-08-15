package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

// -----------------------------------------------------------------------------
type Dog struct {


}

// satisfy the Animal interface
func (dog Dog) Speak() string {
	return "Woof!"
}

// -----------------------------------------------------------------------------
type Cat struct {

}

// satisfy the Animal interface
func (cat Cat) Speak() string {
	return "Meow!"
}

// -----------------------------------------------------------------------------

type Llama struct {

}

// satisfy the Animal interface
func (llama Llama) Speak() string {
	return "?????"
}

// -----------------------------------------------------------------------------
type JavaProgrammer struct {

}

// satisfy the Animal interface
func (j JavaProgrammer) Speak() string {
	return "Design Patterns!"
}
// -----------------------------------------------------------------------------


func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

