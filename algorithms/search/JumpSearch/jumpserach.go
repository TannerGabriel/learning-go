package JumpSearch

import "math"

func jumpSearch(arr []int, query int) int {
	size := len(arr)
	step := int(math.Sqrt(float64(size)))
	prev := 0

	// We can't find anything in empty array
	if size == 0 {
		return -1
	}

	// Finding the block where element is present (if it is present)
	for arr[int(math.Min(float64(step), float64(size)))-1] < query {
		prev = step
		step += int(math.Sqrt(float64(size)))
		if prev >= size {
			return -1
		}
	}

	// Doing a linear search for x in block beginning with prev
	for arr[prev] < query {
		prev++

		// If we reached next block or end of array, element is not present.
		if prev == int(math.Min(float64(step), float64(size))) {
			return -1
		}
	}

	// Check if the element is found
	if arr[prev] == query {
		return prev
	}

	return -1
}
