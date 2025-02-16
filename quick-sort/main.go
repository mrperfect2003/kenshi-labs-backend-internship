package main

import "fmt"

// quickSort recursively sorts the slice using the Quick Sort algorithm.
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr // Base case: arrays with 0 or 1 element are already sorted
	}

	// Choose the pivot as the first element
	pivot := arr[0]

	// Slices to hold elements less than and greater than the pivot
	less := []int{}
	greater := []int{}

	// Partition the array into less and greater slices
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	// Recursively sort the less and greater slices, and concatenate the results
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
	// Example slice
	numbers := []int{33, 10, 55, 71, 29, 5, 70}

	fmt.Println("Original slice:", numbers)

	sortedNumbers := quickSort(numbers)

	fmt.Println("Sorted slice:", sortedNumbers)
}
