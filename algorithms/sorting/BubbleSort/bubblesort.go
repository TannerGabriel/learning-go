package BubbleSort

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}

	return arr
}

