package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{id, name, vacation}
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)

	e2 := Employee{id: 1, name: "David", vacation: false}

	fmt.Printf("%v\n", e2)

	e3 := new(Employee)

	fmt.Printf("%v\n", *e3)

	e4 := NewEmployee(2, "Ernesto", false)

	fmt.Printf("%v\n", *e4)
}
