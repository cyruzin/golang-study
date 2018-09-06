package main

import (
	"fmt"
)

func main() {

	f, l := nameAndAge("Cyro", "Dubeux")

	fmt.Printf("%s %s\n", f, l)

}

func nameAndAge(f string, l string) (string, string) {

	firstName := fmt.Sprintf("Your first name is %s.", f)
	lastName := fmt.Sprintf("Your last name is %s.", l)

	return firstName, lastName
}
