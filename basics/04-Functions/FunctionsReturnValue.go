package main

import "fmt"

// Returning a single value of type int
func add(x int, y int) int {
	return x + y
}

// Named return value
func getArea(l int, b int) (area int) {
	area = l * b
	return // Return without specify variable name
}

// Returning multiple name values
func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return
}

// Passing addresses to a function
func addValue(x *int, y *string) {
	*x = *x + 5
	*y = *y + " World!"
	return
}

func main() {
	// Accepting return value in variable
	sum := add(20, 30)
	fmt.Println("Sum: ", sum)

	// Accepting a named return value
	area := getArea(10, 10)
	fmt.Println("Area: ", area)

	// Accepting multiple return values
	area, parameter := rectangle(10, 10)
	fmt.Println("Area: ", area)
	fmt.Println("Parameter", parameter)

	var number = 20
	var text = "Hello"
	fmt.Println("Before:", text, number)

	addValue(&number, &text)

	fmt.Println("After:", text, number)
}
