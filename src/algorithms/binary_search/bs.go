// Package binarysearch holds a demonstration of the binary search algorithm in Go
package binary_search

func BinarySearch(data []int, target int) int {
	left := 0
	right := len(data) - 1

	for left <= right {
		// This is made to prevent Integer Overflow
		// when we have large slices left and right can become
		// really large numbers that will exceed the maximum value
		// for an integer
		middle := left + (left+right)/2

		elementAtTheMiddle := data[middle]

		if elementAtTheMiddle > target {
			right = middle - 1
		}
		if elementAtTheMiddle < target {
			left = middle + 1
		}
		if elementAtTheMiddle == target {
			return middle
		}
	}
	return 0
}
