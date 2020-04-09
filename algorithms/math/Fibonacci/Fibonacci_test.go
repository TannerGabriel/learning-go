package Fibonacci

import (
	"reflect"
	"testing"
)

func TestRecusiveFibonacci(t *testing.T) {
	t.Run("Fibonacci of 0", func(t *testing.T) {
		got := FibonacciRecursive(0)
		want := 0

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Fibonacci of 5", func(t *testing.T) {
		got := FibonacciRecursive(5)
		want := 5

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Factorial of 8", func(t *testing.T) {
		got := FibonacciRecursive(8)
		want := 21

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Factorial of 10", func(t *testing.T) {
		got := FibonacciRecursive(10)
		want := 55

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestFibonacciSequence(t *testing.T) {
	data := []struct {
		n    int
		want []int
	}{
		{0, []int{0}}, {1, []int{0, 1}}, {2, []int{0, 1, 1}}, {3, []int{0, 1, 1, 2}},
	}

	for _, d := range data {
		if got := fibonacciSequence(d.n); !reflect.DeepEqual(got, d.want) {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}
