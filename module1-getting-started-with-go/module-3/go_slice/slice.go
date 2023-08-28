package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter an integer or 'X' to quit:")
		scanner.Scan()
		input := scanner.Text()

		// If the user enters 'X', exit the program
		if input == "X" {
			break
		}

		// Convert the input to an integer
		num, err := strconv.Atoi(input)
		if err != nil {
			// If the input is not an integer, print an error message and continue
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		// Add the integer to the slice and sort it
		slice = append(slice, num)
		sort.Ints(slice)

		// Print the sorted slice
		fmt.Printf("Sorted slice: %v\n", slice)
	}
}
