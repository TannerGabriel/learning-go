package IsPowerOfTwo

/**
Bitwise represenations of PowerOfTwo numbers always have exactly 1 bit set. The only exception is with a signed integer.
e.g. An 8-bit signed integer with a value of -128 looks like: 10000000

If the number is greater 0 we can check the PowerOfTwo using the bitwise AND operator
*/
func isPowerOfTwoBitwise(num int) bool {
	// Check if num is greater 0
	if num < 0 {
		return false
	}

	return (num & (num - 1)) == 0
}
