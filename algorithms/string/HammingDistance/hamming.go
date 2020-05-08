package HammingDistance

import "log"

func hammingDistance(a, b string) int {
	if len(a) != len(b) {
		log.Fatal("Strings are of different length")
	}

	var distance int = 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance
}
