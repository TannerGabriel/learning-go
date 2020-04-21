package FastPowering

import (
	"testing"
)

func TestFastPowering(t *testing.T) {
	data := []struct {
		n    float64
		k    int
		want float64
	}{
		{1, 0, 1}, {2, 2, 4}, {2, 10, 1024}, {3, 3, 27}, {10, 0, 1}, {16, 16, 18446744073709552000},
	}

	for _, d := range data {
		if got := fastPowering(d.n, d.k); got != d.want {
			t.Errorf("Invalid value for N: %f, got: %f, want: %f", d.n, got, d.want)
		}
	}
}
