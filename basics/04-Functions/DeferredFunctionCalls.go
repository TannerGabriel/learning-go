package main

import "fmt"

func first() {
	fmt.Println("First Function")
}

func second() {
	fmt.Println("Second Function")
}

// Defer is a spezial statement that schedules functions to be executed after the function completes
func main() {
	// The Second function will be called after the first
	defer second()
	first()
}
