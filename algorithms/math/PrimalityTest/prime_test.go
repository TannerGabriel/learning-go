package PrimalityTest

import (
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	data := []struct {
		n    int
		want bool
	}{
		{3, true}, {10, false}, {5, true}, {1024, false}, {29, true}, {7, true}, {11, true},
	}

	for _, d := range data {
		if got := isPrimeNumber(d.n); got != d.want {
			t.Errorf("Invalid value for N: %d, got: %t, want: %t", d.n, got, d.want)
		}
	}
}

func TestIsPrimeNumberUsingSqrt(t *testing.T) {
	data := []struct {
		n    int
		want bool
	}{
		{3, true}, {10, false}, {5, true}, {1024, false}, {29, true}, {7, true}, {11, true},
	}

	for _, d := range data {
		if got := isPrime(d.n); got != d.want {
			t.Errorf("Invalid value for N: %d, got: %t, want: %t", d.n, got, d.want)
		}
	}
}
