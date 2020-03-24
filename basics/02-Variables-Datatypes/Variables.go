package main

import "fmt"

func main() {
	// Declaring variables
	var i int
	var s string

	// Initializing Variables
	i = 20
	s = "Some String"

	fmt.Println(i)
	fmt.Println(s)

	// Creating and initializing variables
	var k int = 35

	fmt.Println(k)

	//Short variable declaration
	j := 50

	fmt.Println(j)

	// Declaring multiple variables
	firstName, lastName := "FirstName", "LastName"

	fmt.Println(firstName + lastName)

	// Variable Declaration Block
	var (
		name = "Donald Duck"
		age  = 50
	)

	fmt.Println(name)
	fmt.Println(age)
}
