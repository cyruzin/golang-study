package main

import (
	"fmt"
)

// Greeting interface
type Greeting interface {
	sayHello() string
}

// Portuguese struct
type Portuguese struct {
	greet string
}

// English struct
type English struct {
	greet string
}

func main() {

	g1 := Portuguese{greet: "Oi"}
	g2 := English{greet: "Hello"}

	printGreeting(g1)
	printGreeting(g2)

}

func (p Portuguese) sayHello() string {
	return p.greet
}

func (e English) sayHello() string {
	return e.greet
}

func printGreeting(g Greeting) {
	fmt.Println(g.sayHello())
}
