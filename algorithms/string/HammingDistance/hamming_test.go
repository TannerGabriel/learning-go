package HammingDistance

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	data := []struct {
		n    string
		k    string
		want int
	}{
		{"a", "a", 0}, {"abc", "azf", 2}, {"1011101", "1001001", 2},
	}

	for _, d := range data {
		if got := hammingDistance(d.n, d.k); got != d.want {
			t.Errorf("Invalid value for N: %v, got: %v, want: %d", d.n, got, d.want)
		}
	}
}
