package main

import "fmt"

func main() {
	hello("Go")
	add(20, 30)
}

func hello(x string) {
	fmt.Printf("Hello %s\n", x)
}

func add(x int, y int) {
	fmt.Println(x + y)
}
