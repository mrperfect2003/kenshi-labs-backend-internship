package main

import "fmt"

// merge combines two sorted slices into a single sorted slice.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Merge the two slices into result
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from left
	result = append(result, left[i:]...)
	// Append any remaining elements from right
	result = append(result, right[j:]...)

	return result
}

// mergeSort recursively divides and sorts the slice.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	// Example slice
	numbers := []int{38, 27, 43, 3, 9, 82, 10}

	fmt.Println("Original slice:", numbers)

	sortedNumbers := mergeSort(numbers)

	fmt.Println("Sorted slice:", sortedNumbers)
}
