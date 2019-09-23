package main

import "log"

func main() {
	closure := func(a string) {
		log.Print(a)
	}

	log.Print("Hello, ")
	closure("World!")
}
