package Factorial


func FactorialRecursive(num int) int {
	if num == 0 {
		return 1
	}
	if num < 0 {
		return num * FactorialRecursive(num+1)
	} else {
		return num * FactorialRecursive(num-1)
	}

}
