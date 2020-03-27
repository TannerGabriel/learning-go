package main

import (
	"fmt"
	"reflect"
)

func main() {
	printMultipleStrings("Hello", "World", "!")
	printMultipleVariables(1, "green", false, 1.314, []string{"foo", "bar", "baz"})
}

// Passing multiple atributes using a variadic function
func printMultipleStrings(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

// Pass multiple different datatypes
func printMultipleVariables(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}
