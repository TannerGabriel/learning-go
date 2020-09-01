package PascalTriangle

import (
	"testing"
)

func TestIsPascalTriangle(t *testing.T) {
	t.Run("PascalTraingle of 1", func(t *testing.T) {
		got := pascalTriangle(1)
		want := [][]int{
			{1},
		}

		if len(got) != len(want) {
			t.Errorf("Different array size then expected")
		}

		for i := 0; i < len(got); i++ {
			if !Equal(got[i], want[i]) {
				t.Errorf("Error")
			}
		}

	})

	t.Run("PascalTraingle of 2", func(t *testing.T) {
		got := pascalTriangle(2)
		want := [][]int{
			{1},
			{1, 1},
		}

		if len(got) != len(want) {
			t.Errorf("Different array size then expected")
		}

		for i := 0; i < len(got); i++ {
			if !Equal(got[i], want[i]) {
				t.Errorf("Error")
			}
		}

	})

	t.Run("PascalTraingle of 5", func(t *testing.T) {
		got := pascalTriangle(5)
		want := [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}

		if len(got) != len(want) {
			t.Errorf("Different array size then expected")
		}

		for i := 0; i < len(got); i++ {
			if !Equal(got[i], want[i]) {
				t.Errorf("Error")
			}
		}

	})
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
