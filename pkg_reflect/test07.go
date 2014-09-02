package main

import (
	"fmt"
	"reflect"
)

type Model interface {
	m()
}

type Company struct{}

func (Company) m() {
	// do stuff
}

type Department struct{}

func (*Department) m() {
	// do stuff
}

type User struct {
	CompanyA    Company
	CompanyB    *Company
	DepartmentA Department
	DepartmentB *Department
}

func (User) m() {
	// do stuff
}

func main() {
	HasModels(&User{})
}

func HasModels(m Model) {
	s := reflect.ValueOf(m).Elem()
	t := s.Type()
	modelType := reflect.TypeOf((*Model)(nil)).Elem()

	for i := 0; i < s.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%d: %s %s -> %t\n", i, f.Name, f.Type, f.Type.Implements(modelType))
	}
}
