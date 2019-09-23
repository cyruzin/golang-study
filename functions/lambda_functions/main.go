package main

import "log"

func main() {
	defer func(a string) {
		log.Print(a)
	}("World!")

	log.Print("Hello, ")
}
