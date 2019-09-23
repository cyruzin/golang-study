package main

import "log"

func main() {
	log.Println(highOrderFunction(5, 5, sum))
	log.Println(highOrderFunction2(hello))
}

func sum(a, b int) int {
	return a + b
}

func hello() string {
	return "Hello"
}

func highOrderFunction(a, b int, f func(a, b int) int) int {
	return f(a, b)
}

func highOrderFunction2(f func() string) string {
	return f() + ", World!"
}
