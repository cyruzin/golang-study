package main

import "log"

func main() {
	log.Println(firstClassFunction(hello)())
}

func hello() string {
	return "Hello, "
}

func firstClassFunction(f func() string) func() string {
	return func() string {
		return f() + "World!"
	}
}
