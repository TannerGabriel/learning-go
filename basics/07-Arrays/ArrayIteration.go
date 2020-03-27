package main

import "fmt"

func main() {
	x := [5]int{0, 5, 10, 15, 20}

	// Standard for loop
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// Getting index and value using range
	for index, element := range x {
		fmt.Println(index, "=>", element)
	}

	// Only getting value using range
	for _, value := range x {
		fmt.Println(value)
	}

	// Range and counter
	j := 0
	for range x {
		fmt.Println(x[j])
		j++
	}
}
