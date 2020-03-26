package main

import (
	"fmt"
)

func main() {
	// If Statement
	x := true
	if x {
		fmt.Println("True")
	}

	// If-Else Statement
	y := 100
	if y > 80 {
		fmt.Println("Greater than 80")
	} else {
		fmt.Println("Lesser than 80")
	}

	// If-Elseif Statement
	grade := 5
	if grade == 1 {
		fmt.Println("You have an A")
	} else if grade > 1 && grade < 5 {
		fmt.Println("You have no A but you are positiv")
	} else {
		fmt.Println("Your grade is negativ")
	}

	// If statement initialization
	if a := 10; a == 10 {
		fmt.Println(a)
	}
}
