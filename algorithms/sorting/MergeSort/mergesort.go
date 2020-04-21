package MergeSort

func mergeSort(arr []int) []int {
	// No need to sort the array if the array only has one element or empty
	if len(arr) <= 1 {
		return arr
	}

	// In order to divide the array in half, we need to figure out the middle
	middle := len(arr) / 2

	// Divide the arrays in a left and right part
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])

	// Combine the left and the right
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	// Concatente values to the result array
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}

		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}
