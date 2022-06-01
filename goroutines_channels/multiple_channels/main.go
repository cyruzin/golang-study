package main

import "fmt"

func main() {
	age := []int{18, 34, 89, 24, 30, 22, 16, 10}

	sum := make(chan int)
	avg := make(chan int)

	go sumAge(sum, age)
	go averageAge(avg, sum, len(age))

	printResult(avg)
}

func sumAge(out chan<- int, age []int) {
	var total int

	for _, v := range age {
		total += v
	}

	out <- total

	close(out)
}

func averageAge(out chan<- int, in <-chan int, lenght int) {
	var average int

	for v := range in {
		average = v / length
		out <- average
	}

	close(out)
}

func printResult(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
