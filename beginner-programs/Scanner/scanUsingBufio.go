package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print(`Which example do you want to run?
  1) fmt.Scan(...)
  2) bufio.Reader.ReadString(...)
Please enter 1 or 2 and press ENTER: `)

	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch result {
	case '1':
		runScan()
		break
	case '2':
		runReadString()
		break
	default:
		return
	}

}

func runScan() {
	// You must declare a variable and then pass the pointer into Scan() function below
	var input string

	fmt.Print("\nPlease insert a string and press ENTER: ")

	// Using the fmt.Scan function, we can read single words as an ASCII string
	num, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(input)
	fmt.Println(num)
}

const INPUTDELIMITER = '\n'

func runReadString() {
	fmt.Print("\nPlease insert a string and press ENTER: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(INPUTDELIMITER)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)

	fmt.Println(input)
	fmt.Println("Exiting program.")
}
