package main

import "fmt"

// linearSearch function searches for the target value in the slice.
// It returns the index of the target if found; otherwise, it returns -1.
func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i // Target found at index i
		}
	}
	return -1 // Target not found
}

func main() {
	// Example slice
	numbers := []int{5, 3, 8, 4, 2}

	// Target value to search for
	target := 4

	// Perform linear search
	index := linearSearch(numbers, target)

	// Output the result
	if index != -1 {
		fmt.Printf("Element %d found at index %d.\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the slice.\n", target)
	}
}
