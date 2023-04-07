package main

import "fmt"

type Person struct {
	name string
	edad int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

type TemporaryEmployee struct {
	Person
	Employee
}

func (person Person) getMessage() string {
	return "Hola mi nombre es " + person.name + "y soy una persona"
}

func (person FullTimeEmployee) getMessage() string {
	return "Hola mi nombre es " + person.name + "y soy un fte"
}

func (person TemporaryEmployee) getMessage() string {
	return "Hola mi nombre es " + person.name + "y soy una te"
}

type PrintInfo interface {
	getMessage() string
}

func GetMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	fte := FullTimeEmployee{}

	fte.name = "David"
	fte.id = 1

	fmt.Printf("%v\n", fte)

	te := TemporaryEmployee{}

	GetMessage(fte)
	GetMessage(te)

}
