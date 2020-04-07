package Factorial

func FactorialRecursive(num int) int {
	if num == 0 {
		return 1
	}

	return num * FactorialRecursive(num-1)
}
