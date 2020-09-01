package SquareRoot

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	data := []struct {
		n    float64
		want float64
	}{
		{4.0, 2.0}, {9.0, 3.0}, {36.0, 6.0}, {81.0, 9.0}, {256.0, 16.0},
	}

	for _, d := range data {
		if got := squareRoot(d.n); got != d.want {
			t.Errorf("Invalid value for N: %f, got: %f, want: %f", d.n, got, d.want)
		}
	}
}
