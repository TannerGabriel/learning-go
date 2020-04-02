package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	path := "test"

	// Create an empty directory
	if _, err := os.Stat(path); os.IsNotExist(err) {
		mdErr := os.Mkdir(path, 0755)
		if mdErr != nil {
			fmt.Println("Error making directory", mdErr)
		}
	}

	// Create file in the directory
	emptyFile, err := os.Create("./test/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emptyFile)
	emptyFile.Close()

	// Read directory recursively
	filepath.Walk(path, func(fn string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error while walking: ", err)
			return err
		}

		if fi.IsDir() {
			fmt.Println("Directory: ", fn)
			if fi.Name() == "skipme" {
				return filepath.SkipDir
			}
		} else {
			fmt.Println("File: ", fn)
		}
		return nil
	})

	// Delete directory
	os.RemoveAll("./test")

	// Get users home directory
	usr, err := user.Current()
	fmt.Println("User Home Directory:", usr.HomeDir)
}
