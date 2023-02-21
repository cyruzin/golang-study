package main

import "fmt"

func minusOne(value int) int {
	if value <= 0 {
		fmt.Println("enter a value greater than 0")
		return 0
	}

	result := value - 1

	fmt.Printf("%d minus 1 is %d\n", value, result)

	if result > 0 {
		return minusOne(result)
	}

	return 0
}

func plusTen(value int) int {
	if value <= 0 {
		fmt.Println("enter a value greater than 0")
		return 0
	}

	if value > 500 {
		fmt.Println("enter a value less than 500")
		return 0
	}

	result := value + 10

	fmt.Printf("%d plus 10 is %d\n", value, result)

	if result < 500 {
		return plusTen(result)
	}

	return 0
}

func main() {
	minusOne(10)

	plusTen(200)
}
