package main

import "fmt"

func main() {
	// Basic while loop
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Do while loop
	num := 0
	for {
		// Work
		fmt.Println(num)

		if num == 10 {
			break
		}
		num++
	}
}
