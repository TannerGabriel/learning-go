package EuclideanAlgorithm

import (
	"testing"
)

func TestEuclideanAlgorithm(t *testing.T) {
	data := []struct {
		n    int
		k    int
		want int
	}{
		{0, 0, 0}, {2, 0, 2}, {0, 2, 2}, {12, 4, 4}, {252, 105, 21}, {-462, -1071, -21},
	}

	for _, d := range data {
		if got := GCD(d.n, d.k); got != d.want {
			t.Errorf("Invalid value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}
