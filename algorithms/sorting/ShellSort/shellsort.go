package ShellSort

func shellSort(array []int) []int {
	length := len(array)

	// Copy the array so the passed one isn't affected
	arr := make([]int, length)
	copy(arr, array)

	// Start with a big gap and reduce it on every iteration
	for gap := int(length / 2); gap > 0; gap /= 2 {
		// Insertion sort for the gap
		for i := gap; i < length; i++ {
			// Shift elements from the gap
			for j := i; j >= gap && arr[j-gap] > arr[j]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}

	return arr
}
