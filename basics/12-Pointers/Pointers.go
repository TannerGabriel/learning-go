package main

import "fmt"

func main() {
	i, j := 10, 1000

	// Point to i
	p := &i
	fmt.Println(*p)

	// Set i using the pointer
	*p = 21
	fmt.Println(i)

	// Point to j
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
