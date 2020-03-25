package main

import "fmt"

// Defining a anonymous function
var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	area := area(10, 10)
	fmt.Println(area)
}
