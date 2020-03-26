package main

import (
	"fmt"
)

func main() {
	// Switch Statement
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Many")
	}

	// Switch Statement with multiple cases
	switch num {
	case 1, 2, 3, 4, 5:
		fmt.Println("Some")
	case 6, 7, 8, 9:
		fmt.Println("More")
	default:
		fmt.Println("Many")
	}

	// Switch fallthrough case statement (forces the execution of the next statement)
	dayOfWeek := 3
	switch dayOfWeek {
	case 1:
		fmt.Println("Go to work")
		fallthrough
	case 2:
		fmt.Println("Buy some bread")
		fallthrough
	case 3:
		fmt.Println("Visit a friend")
		fallthrough
	case 4:
		fmt.Println("Buy some food")
		fallthrough
	case 5:
		fmt.Println("See your family")
	default:
		fmt.Println("No information available for that day.")
	}

	// Switch using conditional statements
	switch {
	case num < 5:
		fmt.Println("Smaller than 5")
	case num == 5:
		fmt.Println("Five")
	case num > 5:
		fmt.Println("Bigger than five")
	default:
		fmt.Println("No information about the number")
	}

	// Switch initializer statement
	switch number := 5; {
	case number < 5:
		fmt.Println("Smaller than 5")
	case number == 5:
		fmt.Println("Five")
	case number > 5:
		fmt.Println("Bigger than five")
	default:
		fmt.Println("No information about the number")
	}
}
