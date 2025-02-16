package main

import "fmt"

// binarySearch function searches for target in the sorted slice arr.
// It returns the index of target if found; otherwise, it returns -1.
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Calculate the middle index'

		if arr[mid] == target { // Target found
			return mid // Target found at index mid
		} else if arr[mid] < target {
			left = mid + 1 // Continue search in the right half
		} else {
			right = mid - 1 // Continue search in the left half
		}
	}

	return -1 // Target not found
}

func main() {
	// Sorted slice of integers
	numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	// Target value to search for
	target := 14

	// Perform binary search
	index := binarySearch(numbers, target)

	// Output the result
	if index != -1 {
		fmt.Printf("Element %d found at index %d.\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the slice.\n", target)
	}
}
