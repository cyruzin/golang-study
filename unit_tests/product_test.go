package main

import "testing"

func Product(x int, y int) int {
	return x * y
}

func TestProduct(t *testing.T) {
	// Product of positive numbers
	test1 := Product(4, 5)
	if test1 != 20 {
		t.Errorf("Product(4,5) = %d; expected 20", test1)
	}

	// Product of a positive number and 0
	test2 := Product(4, 0)
	if test2 != 0 {
		t.Errorf("Product(4,0) = %d; expected 0", test2)
	}

	// Product of a positive and a negative number
	test3 := Product(4, -5)
	if test3 != -20 {
		t.Errorf("Product(4,-5) = %d; expected -20", test3)
	}
}
