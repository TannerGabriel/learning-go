package main

import (
	"fmt"
	"reflect"
)

// Slices are dynamic datastructures that can grow and shrink as you see fit.
func main() {
	// Create an empty slice
	var x []int
	fmt.Println(reflect.ValueOf(x).Kind())

	// Creating a slice using the make function
	var y = make([]string, 10, 20)
	fmt.Printf("y \tLen: %v \tCap: %v\n", len(y), cap(y))

	// Initialize the slice with values using a slice literal
	var z = []int{10, 20, 30, 40}
	fmt.Printf("z \tLen: %v \tCap: %v\n", len(z), cap(z))
	fmt.Println(z)

	// Creating a Slice using the new keyword
	var a = new([50]int)[0:10]
	fmt.Printf("a \tLen: %v \tCap: %v\n", len(a), cap(a))
	fmt.Println(a)

	// Add items using the append function
	var b = make([]int, 1, 10)
	fmt.Println(b)
	b = append(b, 20)
	fmt.Println(b)

	// Access slice items
	var c = []int{10, 20, 30, 40}
	fmt.Println(c[0])
	fmt.Println(c[0:3])

	// Change item values
	var d = []int{10, 20, 30, 40}
	fmt.Println(d)
	d[1] = 35
	fmt.Println(d)

	// Copy slice into another slice
	var e = []int{10, 20, 30, 40}
	var f = []int{50, 60, 70, 80}
	copy(e, f)
	fmt.Println("E: ", e)

	// Append a slice to an existing one
	var g = []int{10, 20, 30, 40}
	var h = []int{50, 60, 70, 80}

	g = append(g, h...)
	fmt.Println(g)
}
