package main

import "fmt"

// Declaring constants
const PI float64 = 3.14159265359
const VALUE = 1000

// Multilple Constants Declaration Block
const (
	PRODUCT  = "Ice Cream"
	QUANTITY = 50
)

func main() {
	fmt.Println(PI)
	fmt.Println(VALUE)

	fmt.Println(PRODUCT)
	fmt.Println(QUANTITY)
}
