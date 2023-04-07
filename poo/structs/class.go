package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	employee := Employee{}
	fmt.Printf("%v", employee)
	employee.id = 0
	employee.name = "John Doe"
	fmt.Printf("%v", employee)
	employee.SetId(2)
	employee.SetName("David")
	fmt.Printf("%v", employee)
	// Print with getters
	fmt.Println(employee.GetName())
	fmt.Println(employee.GetId())

}
