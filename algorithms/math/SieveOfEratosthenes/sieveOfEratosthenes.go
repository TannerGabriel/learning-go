package SieveOfEratosthenes

func sieveOfEratosthenes(maxNumber int) []int {
	isPrime := make([]bool, maxNumber+1)

	// Initialize array with true values
	for i := 0; i < maxNumber+1; i++ {
		isPrime[i] = true
	}

	isPrime[0] = false
	isPrime[1] = false

	var primes []int

	// Loop until max number
	for num := 2; num <= maxNumber; num++ {
		if isPrime[num] == true {
			// Save prime number in array
			primes = append(primes, num)

			// Start marking multiples of p from p * p, and not from 2 * p because, smaller multiples of p will have already been marked false.
			nextNum := num * num

			// Mark all multiples of a prime number as false
			for nextNum <= maxNumber {
				isPrime[nextNum] = false
				nextNum += num
			}
		}
	}

	return primes
}
