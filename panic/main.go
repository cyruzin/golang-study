// Panic means something went unexpectedly wrong.
// We use panic to fail fast on errors that shouldn’t occur normally or weren't handles gracfully

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Enter a number less than 100: ")
	var input string
	fmt.Scanln(&input)
	intNum, err := stringToInt(input)
	if err != nil {
		panic("Not a number")
	}
	if intNum > 100 {
		panic("Don't know how to handle that!")
	}
	fmt.Printf("You entered the number %d \n", intNum)
}

func stringToInt(n string) (int, error) {
	num, err := strconv.Atoi(n)
	var re int
	if err != nil {
		return re, err
	}
	return num, nil
}
