package main

import "fmt"

func main() {
	var mf myFloat = 0.5

	fmt.Println(myFunc((mf)))
}

type myFloat float64

// In cases like this, you need the tilde sign
func myFunc[T ~float64](a T) T {
	return a
}
