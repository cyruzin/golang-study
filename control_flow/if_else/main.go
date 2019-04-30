package main

import "fmt"

func main() {
	// basic
	var x = 25

	if x%5 == 0 {
		fmt.Printf("%d is a multiple of 5\n", x)
	}

	// short
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	}

	// chain
	var age = 33

	if age < 32 {
		fmt.Println("Your age is less than 32")
	} else if age == 32 {
		fmt.Println("Your age is 32")
	} else {
		fmt.Println("Your age is greater than 32")
	}

}
