package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 uint
}

func (p *person) listPerson() {
	fmt.Printf("Hello, %s %s.\n", p.firstName, p.lastName)
	fmt.Printf("You are %d years old.\n", p.age)
}

func main() {
	// To access a receiver you need to instantiate its type
	p := &person{
		firstName: "Cyro",
		lastName:  "Dubeux",
		age:       32,
	}

	p.listPerson()
}
