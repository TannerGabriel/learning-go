package IsPowerOfTwo

import (
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	data := []struct {
		n    int
		want bool
	}{
		{2, true}, {10, false}, {64, true}, {1024, true},
	}

	for _, d := range data {
		if got := isPowerOfTwo(d.n); got != d.want {
			t.Errorf("Invalid value for N: %d, got: %t, want: %t", d.n, got, d.want)
		}
	}
}

func TestIsPowerOfTwoBitwise(t *testing.T) {
	data := []struct {
		n    int
		want bool
	}{
		{2, true}, {10, false}, {64, true}, {1024, true},
	}

	for _, d := range data {
		if got := isPowerOfTwoBitwise(d.n); got != d.want {
			t.Errorf("Invalid value for N: %d, got: %t, want: %t", d.n, got, d.want)
		}
	}
}
