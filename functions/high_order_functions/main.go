package main

import "log"

func main() {
	log.Println(highOrderFunction(hello)())
}

func hello() string {
	return "Hello, "
}

func highOrderFunction(f func() string) func() string {
	return func() string {
		return f() + "World!"
	}
}
