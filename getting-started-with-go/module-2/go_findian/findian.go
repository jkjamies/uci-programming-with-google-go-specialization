package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the newline character
	input = strings.TrimSpace(input)

	if containsIan(input) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func containsIan(s string) bool {
	cleanedInput := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return strings.Contains(cleanedInput, "i") &&
		strings.Contains(cleanedInput, "a") &&
		strings.Contains(cleanedInput, "n")
}
