package main

import "log"

func main() {
	log.Println(firstClassFunction(5, 5, sum))
	log.Println(firstClassFunction2(hello))
}

func sum(a, b int) int {
	return a + b
}

func hello() string {
	return "Hello"
}

func firstClassFunction(a, b int, f func(a, b int) int) int {
	return f(a, b)
}

func firstClassFunction2(f func() string) string {
	return f() + ", World!"
}
