package main

import "fmt"

// Divide and conquer!
//
// Binary search works on sorted arrays. Binary search begins by comparing
// the middle element of the array with the target value. If the target
// value matches the middle element, its position in the array is returned.
// If the target value is less than the middle element, the search continues
// in the lower half of the array. If the target value is greater than the
// middle element, the search continues in the upper half of the array.
// By doing this, the algorithm eliminates the half in which the target
// value cannot lie in each iteration.
func binarySearch(key int, vector []int) int {
	var inferior int = 0
	var superior int = len(vector) - 1
	var middle int

	for inferior <= superior {
		middle = (inferior + superior) / 2

		if key == vector[middle] {
			return middle
		}

		if key < vector[middle] {
			superior = middle - 1
		} else {
			inferior = middle + 1
		}
	}

	return -1
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	position := binarySearch(632, items)

	if position > 0 {
		fmt.Println("Key found at position:", position)
	} else {
		fmt.Println("Key not found!")
	}
}
