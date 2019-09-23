package main

import "fmt"

func main() {
	fmt.Println("The sum is: ", sum(5, 5))
	sayHello()
}

// void function
func sayHello() {
	fmt.Println("Hello :)")
}

// function with return
func sum(x, y float64) float64 {
	return x + y
}
