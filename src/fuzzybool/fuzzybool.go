package main

import (
	"fmt"
)

// -----------------------------------------------------------------------------
// 								FUZZYBOOL
// -----------------------------------------------------------------------------
type FuzzyBool struct {
	value float32
}

// constructor, returns a pointer that points to a value of type FuzzyBool
func New(value interface{}) (*FuzzyBool, error) {
	amount, err := float32ForValue(value)
	return &FuzzyBool{amount}, err
}

func float32ForValue(value interface{}) (fuzzy float32, err error) {
	switch value := value.(type) { // shadow variable
	case float32:
		fuzzy = value
	case float64:
		fuzzy = float32(value)
	case int:
		fuzzy = float32(value)
	case bool:
		fuzzy = 0
		if value {
			fuzzy = 1
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

// copy constructor
func (fuzzy *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool{fuzzy.value}
}

// satisfy interface `fmt.Stringer`
func (fuzzy *FuzzyBool) String() string {
	return fmt.Sprintf("%3.0f%%", 100*fuzzy.value)
}

func (fuzzy *FuzzyBool) Set(value interface{}) (err error) {
	fuzzy.value, err = float32ForValue(value)
	return err
}

func (fuzzy *FuzzyBool) Not() *FuzzyBool {
	return &FuzzyBool{1 - fuzzy.value}
}

func (fuzzy *FuzzyBool) And(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	minimum := fuzzy.value
	rest = append(rest, first)

	for _, other := range rest {
		if minimum > other.value {
			minimum = other.value
		}
	}
	return &FuzzyBool{minimum}
}

func (fuzzy *FuzzyBool) Or (first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	maximum := fuzzy.value
	rest = append(rest, first)

	for _, other := range rest {
		if maximum < other.value {
			maximum = other.value
		}
	}
	return &FuzzyBool{maximum}
}

func (fuzzy *FuzzyBool) Less(other *FuzzyBool) bool {
	return fuzzy.value < other.value
}

func (fuzzy *FuzzyBool) Greater(other *FuzzyBool) bool {
	return fuzzy.value > other.value
}

func (fuzzy *FuzzyBool) Equal(other *FuzzyBool) bool {
	return fuzzy.value == other.value
}

func (fuzzy *FuzzyBool) Bool() bool {
	return fuzzy.value > 0.5
}

func (fuzzy *FuzzyBool) Float() float64 {
	return float64(fuzzy.value)
}

// -----------------------------------------------------------------------------
func main() {
	a, _ := New(0)		// can safely ignore the value of err
	b, _ := New(.25)		// confirmed as valid value, need confirm when use
	c, _ := New(.75)		// still is a variable
	d    := c.Copy()

	if err :=d.Set(1); err != nil {
		fmt.Println(err)
	}

	process(a, b, c, d)

	s := []*FuzzyBool{a, b, c, d}
	fmt.Println(s)
}

func process(a, b, c, d *FuzzyBool) {
	fmt.Println("Original: ", a, b, c, d)
	fmt.Println("Not     : ", a.Not(), b.Not(), c.Not(), d.Not())
	fmt.Println("Not Not : ", a.Not().Not(), b.Not().Not(), c.Not().Not(), d.Not().Not())

	fmt.Println("0.And(.25)     : →", a.And(b))
	fmt.Println(".25.And(.75)   : →", b.And(c))
	fmt.Println(".75.And(1)     : →", c.And(d))
	fmt.Println(".25.And(.75, 1): →", b.And(c, d))

	fmt.Println("0.Or(.25)      : →", a.Or(b))
	fmt.Println(".25.Or(.75)    : →", b.Or(c))
	fmt.Println(".75.Or(1)      : →", c.Or(d))
	fmt.Println(".25.Or(.75, 1) : →", b.Or(c, d))

	fmt.Println("a < c, a == c, a > c:", a.Less(c), a.Equal(c), a.Greater(c))

	fmt.Println("Bool:,,", a.Bool(), b.Bool(), c.Bool(), d.Bool())
	fmt.Println("Float: ", a.Float(), b.Float(), c.Float(), d.Float())
}

