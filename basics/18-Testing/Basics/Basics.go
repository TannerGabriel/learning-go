package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Hello " + name
}

func Add(x, y int) int {
	return x + y
}

func Multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(Hello("Gabriel"))
}
