package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	for {
		truncatedNumber, continueLoop, err := truncUserFloatingPointInput()
		if !continueLoop {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Truncated number:", truncatedNumber)
	}
}

func truncUserFloatingPointInput() (int, bool, error) {
	fmt.Print("Enter a floating point number (or 'exit' to quit): ")

	var inputStr string
	_, err := fmt.Scanln(&inputStr)
	if err != nil {
		return 0, false, err
	}

	if strings.ToLower(inputStr) == "exit" {
		return 0, false, nil
	}

	var input float64
	_, err = fmt.Sscan(inputStr, &input)
	if err != nil {
		return 0, true, fmt.Errorf("not a valid floating point number")
	}

	return int(math.Trunc(input)), true, nil
}
