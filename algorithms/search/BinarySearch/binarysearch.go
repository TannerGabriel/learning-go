package BinarySearch

func binarySearch(arr []int, query int) int {
	// Set min and max index
	minIndex := 0
	maxIndex := len(arr) - 1

	for minIndex <= maxIndex {
		// Calculate the mid index and get the current mid item
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := arr[midIndex]

		// If the item is found return it
		if query == midItem {
			return midIndex
		}

		// Set the new min and max index of the subarray
		if midItem < query {
			minIndex = midIndex + 1
		} else if midItem > query {
			maxIndex = midIndex - 1
		}
	}

	return -1
}
