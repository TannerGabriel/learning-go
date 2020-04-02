package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Create an empty file
	emptyFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emptyFile)
	emptyFile.Close()

	// Rename a file
	oldName := "test.txt"
	newName := "test2.txt"
	os.Rename(oldName, newName)

	// Check if file exists
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("File: test.txt does not exist")
	}

	// Get file information
	fileStat, err := os.Stat("test2.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time and date

	// Writing to a file
	outfile := "test.txt"
	content := []byte("Some great text")
	err = ioutil.WriteFile(outfile, content, 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
	} else {
		fmt.Println("Created: ", outfile)
	}

	// Reading a file
	filename := "test.txt"

	fileContent, fileError := ioutil.ReadFile(filename)
	if fileError != nil {
		log.Fatalln("Error reading file", filename)
	} else {
		fmt.Println(string(fileContent))
	}

	// Append to an existing file
	af, err := os.OpenFile("test2.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error appending to file:", err)
	}
	if _, err = af.WriteString("\nAppending this text"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
	af.Close()

	// Reduce the file size
	err = os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}

	// Delete a specific file
	err = os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
}
