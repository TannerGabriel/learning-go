package main

import "fmt"

func recoveryFunction() {
	// Recover from panic and print erro message
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
	fmt.Println("End: recoveryFunction")
}

func executePanic() {
	// Defer function call
	defer recoveryFunction()
	// Throw panic
	panic("Panic")
	fmt.Println("End: executePanic")
}

func main() {
	executePanic()
	fmt.Println("End: main")
}
