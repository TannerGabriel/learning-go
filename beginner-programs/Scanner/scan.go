package main

import "fmt"

func main() {
	var s string

	// Read a string using fmt.Scan()
	fmt.Print("Please insert a string an press enter: ")
	fmt.Scan(&s)

	// Print the read string
	fmt.Printf("You entered \"%v\"", s)
}
