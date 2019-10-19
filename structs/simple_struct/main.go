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

	fmt.Printf("%+v\n", p)
	fmt.Println("Changing name and lastname...")
	p.changeName("Paul", "Jones")
	fmt.Printf("%+v\n", p)

}

func (p *person) changeName(f, l string) {
	p.firstName = f
	p.lastName = l
}
