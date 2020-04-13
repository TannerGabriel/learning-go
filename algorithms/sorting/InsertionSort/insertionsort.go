package InsertionSort

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				swap(arr, j, i)
			}
		}
	}

	return arr
}

func swap(arr []int, i, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}
