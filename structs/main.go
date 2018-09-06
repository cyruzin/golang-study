package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	email     string
	country   string
	age       int
}

func main() {

	p := person{
		firstName: "Cyro",
		lastName:  "Dubeux",
		email:     "xorycx@gmail.com",
		country:   "Brazil",
		age:       31,
	}

	p.changeName("Pedro", "Alvex")
	p.changeName("Cyro", "Dubeux")

	fmt.Printf("%+v\n", p)

}

func (p *person) changeName(f, l string) {
	(*p).firstName = f
	(*p).lastName = l
}
