package Factorial

func Factorial(num int) int {
	result := 1

	for i := 1; i <= num; i++ {
		result *= i
	}

	return result
}
