package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}

// Function that returns another function
func partialMultiplication(x int) func(int) int {
	return func(y int) int {
		return multiply(x, y)
	}
}

func main() {
	multiple := partialMultiplication(10)
	fmt.Println(multiple(10))
}
