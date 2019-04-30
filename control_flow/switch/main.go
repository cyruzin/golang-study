package main

import "fmt"

func main() {
	// basic
	age := 32

	switch age {
	case 1:
		fmt.Println("You are young")
	case 32:
		fmt.Println("Bullseye!")
	default:
		fmt.Println("How old are you?")
	}

	// short
	switch color := "red"; color {
	case "red":
		fmt.Println("I love the red color")
	case "yellow":
		fmt.Println("I don't like the yellow color")
	case "green":
		fmt.Println("Green is nice")
	default:
		fmt.Println("Black is the best color")
	}
}
