package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt user for name
	fmt.Print("Enter Name: ")
	// Read user's name input
	inputName, _ := reader.ReadString('\n')

	// Prompt user for address
	fmt.Print("Enter Address: ")
	// Read user's address input
	inputAddress, _ := reader.ReadString('\n')

	// Create a map to store the data
	data := make(map[string]string)

	// Store the data in the map, trimming whitespace to avoid printing "\n"
	data["name"] = strings.TrimSpace(inputName)
	data["address"] = strings.TrimSpace(inputAddress)

	// Convert the map to JSON
	jsonData, err := json.Marshal(data)
	// Check for errors and print them, then exit if there are any
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print the JSON
	fmt.Println(string(jsonData))
}
