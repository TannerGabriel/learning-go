package main

import "fmt"

func main() {
	l := 10
	b := 10

	// Closure functions are a special case of a anonymous function where you access outside variables
	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
}
