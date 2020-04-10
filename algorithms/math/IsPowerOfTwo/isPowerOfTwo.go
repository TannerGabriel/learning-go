package IsPowerOfTwo

/*
	Find if the number is a power of two or not by dividing it without remainder (primitive)
*/
func isPowerOfTwo(num int) bool {
	if num < 2 {
		return false
	}

	for num != 1 {
		if mod(num, 2) != 0 {
			return false
		}
		num = num / 2
	}

	return true
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
