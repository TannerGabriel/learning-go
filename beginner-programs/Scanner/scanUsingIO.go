package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Printf("Enter the text:\n")
	writeText, err := os.Open(os.DevNull)

	if err != nil {
		log.Fatalf("failed to open a null device: %s", err)
	}

	defer writeText.Close()
	io.WriteString(writeText, "Write Text")

	readText, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatalf("failed to read stdin: %s", err)
	}

	fmt.Printf("\nLength: %d", len(readText))
	fmt.Printf("\nData Read: \n%s", readText)
}
