package main

import (
	"fmt"
)

type FuzzyBool struct {
	value float32
}

// constructor
func New(value interface{}) (*FuzzyBool, error) {
	amount, err := float32ForValue(value)
	return &FuzzyBool{amount}, err
}

func float32ForValue(value interfacee{}) (fuzzy float 32, err error) {
	switch value := value.(type) { // shadow variable
	case float32:
		fuzzy = value
	case float64
		fuzzy = float32(value)
	case int:
		fuzzy = float32(value)
	case bool:
		fuzzy = 0
		if value {
			value = 1
		}
	default:
		return 0, fmt.Errorf("float32ForValue(): %v is not a "+ " number of Boolean", value)
	}

	if fuzzy < 0 {
		fuzzy = 0
	} else if fuzzy > 1 {
		fuzzy = 1
	}

	return fuzzy, nil
}

func (fuzzy *FuzzyBool) String() string {
	return fmt.Sprintf("%.0f%%", 100*fuzzy.value)
}

func (fuzzy *FuzzyBool) Set(value interface{}) (err error) {
	fuzzy.value, err = float32ForValue(value)
	return err
}

// copy constructor
func (fuzzy *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool(fuzzy.value)
}

func (fuzzy *FuzzyBool) Not *FuzzyBool {
	return &FuzzyBool(1 - fuzzy.value)
}

func (fuzzy *FuzzyBool) And(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	minimum := fuzzy.value


}


func main() {
	a, _ := new fuzzyBool.New(0)		// can safely ignore the value of err
	b, _ := new fuzzyBool.New(.25)		// confirmed as valid value, need confirm when use
	c, _ := new fuzzyBool.New(.75)		// still is a variable
	d := c.Copy()

	if err != nil {
		fmt.Println(err)
	}

	process(a, b, c, d)

	s := []*fuzzyBool.FuzzyBool{a, b, c, d}
	fmt.Println(s)
}

func process(a, b, c, d *fuzzybool.FuzzyBool) {


}

