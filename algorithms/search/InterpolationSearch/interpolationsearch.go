package InterpolationSearch

func interpolationSearch(arr []int, query int) int {
	// Find indexes of two corners
	lo := 0
	hi := len(arr) - 1

	// Since array is sorted, an element present in array must be in range defined by corner
	for lo <= hi && query >= arr[lo] && query <= arr[hi] {
		// Probing the position with keeping uniform distribution in mind.
		midIndex := lo + (query-arr[lo])*(hi-lo)/(arr[hi]-arr[lo])
		midItem := arr[midIndex]

		// Target is found
		if midItem == query {
			return midIndex
		} else if midItem < query {
			// Target is in upper part
			lo = midIndex + 1
		} else if midItem > query {
			// Target is in lower part
			hi = midIndex - 1
		}
	}

	return -1
}
