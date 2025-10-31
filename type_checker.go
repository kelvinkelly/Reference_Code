package main

import (
	"fmt"
	"os"
	"strconv"
)

// Define a map to store descriptions of categories
var categories = map[string]string{
	"numeric": "This value looks like an integer, likely a counter or index.",
	"textual": "This is a string, often used for names or messages.",
	"other":   "This is uncategorized or missing a value.",
}

// checkType examines the input string and returns a category string.
func checkType(input string) string {
	if input == "" {
		return "other"
	}

	// Try to parse as an integer
	_, err := strconv.Atoi(input)
	if err == nil {
		return "numeric"
	}

	// Simple check: if it's not a number, we'll treat it as text
	return "textual"
}

func main() {
	// 1. Get the command-line arguments
	// os.Args[0] is the program name itself.
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run <program_name.go> <some_value>")
		fmt.Printf("Defaulting to 'other' category. %s\n", categories["other"])
		return
	}

	// The argument we care about is the second element (index 1)
	input := os.Args[1]
	
	// 2. Call the checker function
	category := checkType(input)

	// 3. Use a switch statement to handle the result
	fmt.Printf("Input Value: **%s**\n", input)
	fmt.Print("--- Status ---\n")

	switch category {
	case "numeric":
		fmt.Printf("Category: 🔢 **Numeric**\n")
		fmt.Println(categories["numeric"])
	case "textual":
		fmt.Printf("Category: 📝 **Textual**\n")
		fmt.Println(categories["textual"])
	default:
		fmt.Printf("Category: ❓ **Other/Unknown**\n")
		fmt.Println(categories["other"])
	}
}
