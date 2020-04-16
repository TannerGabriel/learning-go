package PrimalityTest

import (
	"math"
)

// O(n)
func isPrimeNumber(num int) bool {
	for i := 2; i < num; i++ {
		if mod(num, i) == 0 {
			return false
		}
	}

	return true
}

// O(sqrt(n))
func isPrime(num int) bool {
	s := math.Sqrt(float64(num))
	for i := 2; i < int(s); i++ {
		if mod(num, i) == 0 {
			return false
		}
	}
	return num > 1
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}
