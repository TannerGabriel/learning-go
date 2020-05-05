package Radian

import (
	"math"
	"testing"
)

func TestDegreeToRadian(t *testing.T) {
	data := []struct {
		n    float64
		want float64
	}{
		{45, math.Pi / 4}, {90, math.Pi / 2}, {180, math.Pi}, {270, 3 * math.Pi / 2},
	}

	for _, d := range data {
		if got := degreeToRadian(d.n); got != d.want {
			t.Errorf("Invalid value for N: %f, got: %f, want: %f", d.n, got, d.want)
		}
	}
}

func TestRadianToDegree(t *testing.T) {
	data := []struct {
		n    float64
		want float64
	}{
		{math.Pi / 4, 45}, {math.Pi / 2, 90}, {math.Pi, 180}, {3 * math.Pi / 2, 270},
	}

	for _, d := range data {
		if got := radianToDegree(d.n); got != d.want {
			t.Errorf("Invalid value for N: %f, got: %f, want: %f", d.n, got, d.want)
		}
	}
}
