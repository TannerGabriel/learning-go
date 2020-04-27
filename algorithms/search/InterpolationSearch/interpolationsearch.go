package InterpolationSearch

func interpolationSearch(arr []int, query int) int {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi && query >= arr[lo] && query <= arr[hi] {
		midIndex := lo + (query-arr[lo])*(hi-lo)/(arr[hi]-arr[lo])
		midItem := arr[midIndex]
		if midItem == query {
			return midIndex
		} else if midItem < query {
			lo = midIndex + 1
		} else if midItem > query {
			hi = midIndex - 1
		}
	}

	return -1
}
