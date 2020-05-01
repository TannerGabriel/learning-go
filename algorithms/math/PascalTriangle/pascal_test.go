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

// expect(pascalTriangle(0)).toEqual([1]);
//     expect(pascalTriangle(1)).toEqual([1, 1]);
//     expect(pascalTriangle(2)).toEqual([1, 2, 1]);
//     expect(pascalTriangle(3)).toEqual([1, 3, 3, 1]);
//     expect(pascalTriangle(4)).toEqual([1, 4, 6, 4, 1]);
//     expect(pascalTriangle(5)).toEqual([1, 5, 10, 10, 5, 1]);
//     expect(pascalTriangle(6)).toEqual([1, 6, 15, 20, 15, 6, 1]);
//     expect(pascalTriangle(7)).toEqual([1, 7, 21, 35, 35, 21, 7, 1]);
