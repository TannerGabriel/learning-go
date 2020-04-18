package SelectionSort

func selectionSort(arr []int) []int {
	array := arr
	var min int = 0

	// Loop through array
	for i := 0; i < len(arr); i++ {
		min = i

		// Loop through subarray
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		array[i], array[min] = array[min], array[i]
	}

	return array
}
