package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Basic Regexp for matching a number
	pattern := "[0-9]+"

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// Test if the pattern works
	str := "Some string 0 1 2 3 4 "
	if re.MatchString(str) {
		fmt.Println("Yes, matched a number")
	} else {
		fmt.Println("No, no match")
	}

	// Return match
	result := re.FindString(str)
	fmt.Println("Number matched:", result)

	// Return multiple matches
	results := re.FindAllString(str, -1)
	fmt.Println("Number matched: ", results)

	// Replace match
	replaceResults := re.ReplaceAllString(str, "num")
	fmt.Println("Result:", replaceResults)

	// Submatches
	str1 := "Hello @world@ Match"
	sub_re, _ := regexp.Compile("@(.*)@")

	m := sub_re.FindStringSubmatch(str1)
	if len(m) > 1 {
		fmt.Println(m[1])
	}
}
