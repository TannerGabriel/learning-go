package Fibonacci

// Saving numbers into a slice
func fibonacciSequence(num int) []int {
	var res = []int{}
	for n := 0; n <= num; n++ {
		result := fibonacci(n)
		res = append(res, result)
	}
	return res
}

// Calculating fibonacci number
func fibonacci(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}
