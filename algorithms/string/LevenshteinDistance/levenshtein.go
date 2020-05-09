package LevenshteinDistance

func levenshteinDistance(a, b string) int {
	// If a string is empty return the length of the second one
	if len(a) == 0 {
		return len(b)
	}

	if len(b) == 0 {
		return len(a)
	}

	// If the strings are the same return 0
	if a == b {
		return 0
	}

	// Swap longer string to b to save some memory O(min(a,b)) instead of O(a)
	if len(a) > len(b) {
		a, b = b, a
	}
	lenS1 := len(a)
	lenS2 := len(b)

	// Declare the array
	x := make([]uint16, lenS1+1)

	// Fill array
	for i := 1; i < len(x); i++ {
		x[i] = uint16(i)
	}

	// Use formula to fill in the rest of the row
	for i := 1; i <= lenS2; i++ {
		prev := uint16(i)
		var current uint16
		for j := 1; j <= lenS1; j++ {
			if b[i-1] == a[j-1] {
				current = x[j-1]
			} else {
				current = min(min(x[j-1]+1, prev+1), x[j]+1)
			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}

	return int(x[lenS1])
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}
