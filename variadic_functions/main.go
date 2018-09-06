package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(10, 7.6, 11.7, 100, 500.3))

}

func sum(nums ...float64) float64 {
	s := 0.0

	for _, n := range nums {
		s += n
	}

	return s
}
