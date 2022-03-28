package main

import "fmt"

func main() {
	myNumbers := []int{1, 2, 3}
	Print(myNumbers)
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
