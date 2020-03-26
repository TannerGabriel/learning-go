package main

import (
	"fmt"
)

func main() {
	// Basic for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// Infinite loop
	i := 0
	for {
		fmt.Println("Hello World!")
		// Breaks/Stops the infinite loop
		if i == 10 {
			break
		}
		i++
	}

	// Ranges
	strings := []string{"Hello", "World", "!"}

	// Get the index and value while looping through the range
	for i, val := range strings {
		fmt.Printf("%d: %s \n", i, val)
	}

	// Only getting the index
	for i := range strings {
		fmt.Println(i)
	}

	// Only getting the value
	for _, val := range strings {
		fmt.Println(val)
	}
}
