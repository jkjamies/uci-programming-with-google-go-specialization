package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// name is a struct that holds a first and last name
type name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	// Prompt the user for a file name
	fmt.Print("Enter the name of the file: ")
	// Read the file name from the user
	_, err := fmt.Scan(&fileName)
	// Check for errors and exit if there is one
	if err != nil {
		return
	}

	// Open the file
	file, err := os.Open(fileName)
	// Check for errors, print error, exit
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// Close the file when we are done
	defer func(file *os.File) {
		err := file.Close()
		// Check for errors, print error, exit
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
	}(file)

	// Read the file line by line
	var names []name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into first and last name
		parts := strings.Split(line, " ")
		// Add the name to the names slice
		if len(parts) >= 2 {
			names = append(names, name{fname: parts[0], lname: parts[1]})
		}
	}

	// Check for errors, print error, exit
	if scanner.Err() != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	// Print the names
	for _, name := range names {
		fmt.Printf("%s %s\n", name.fname, name.lname)
	}
}
