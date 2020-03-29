package main

import "fmt"

var m = map[string]string{
	"c": "Cyan",
	"y": "Yellow",
	"m": "Magenta",
	"k": "Black",
}

func main() {
	// Declaring empty map
	var shopingList = map[string]int{}
	fmt.Println(shopingList)

	// Initializing a map
	var people = map[string]int{"Elon": 10, "Jeff": 15}
	fmt.Println(people)

	// Map declaration using make function
	var peopleList = make(map[string]int)
	peopleList["Elon"] = 10
	peopleList["Jeff"] = 15
	fmt.Println(peopleList)

	// Accessing items
	fmt.Println(m["c"])

	// Adding items
	m["b"] = "black"
	fmt.Println(m)

	// Updating items
	m["y"] = "lemon yellow"
	fmt.Println(m)

	// Deleting items
	delete(m, "b")
	fmt.Println(m)

	// Iterating over a map
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %s", k, v)
	}

	// Test if an item exists
	c, ok := m["y"]
	fmt.Println("\nc: ", c)
	fmt.Println("ok: ", ok)
}
