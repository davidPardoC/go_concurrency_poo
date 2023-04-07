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

func GetMessage(person Person) {
	fmt.Println("Hola mi edad es", person.edad)
}

func main() {
	fte := FullTimeEmployee{}

	fte.name = "David"
	fte.id = 1

	fmt.Printf("%v\n", fte)

	// Does not work
	// GetMessage(fte)
}
