package CountingSort

func getCountArrayLength(arr []int) int {
	// If array is 0 + 1
	if len(arr) == 0 {
		return 1
	}

	k := arr[0]

	// Set k to the highest value of the array
	for _, value := range arr {
		if value > k {
			k = value
		}
	}

	return k + 1
}

func countingSort(arr []int) []int {
	k := getCountArrayLength(arr)
	// Create array with a length of k
	count := make([]int, k)

	// Store count of each character
	for i := 0; i < len(arr); i++ {
		count[arr[i]] += 1
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			if count[i] > 0 {
				arr[j] = i
				j += 1
				count[i] -= 1
				continue
			}
			break
		}
	}

	return arr
}
