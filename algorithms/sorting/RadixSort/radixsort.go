package RadixSort

// Get maximum value of an array
func getMax(arr []int, n int) int {
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

// A function to do counting sort of arr[] according to the digit represented by exp.
func countSort(arr []int, n int, exp int) {
	output := make([]int, n)
	var count = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	i := 0

	// Store count of occurrences in count[]
	for i = 0; i < n; i++ {
		count[(arr[i]/exp)%10]++
	}

	// Change count[i] so that count[i] now contains actual position of this digit in output[]
	for i = 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i = n - 1; i >= 0; i-- {
		output[count[(arr[i]/exp)%10]-1] = arr[i]
		count[(arr[i]/exp)%10]--
	}

	// Copy the output array to arr[], so that arr[] now contains sorted numbers according to curent digit
	for i = 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func radixsort(arr []int, n int) {
	m := getMax(arr, n)

	for exp := 1; m/exp > 0; exp *= 10 {
		countSort(arr, n, exp)
	}
}
