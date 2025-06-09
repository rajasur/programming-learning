package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Example 1: Basic string input
	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove trailing newline
	fmt.Printf("Hello, %s!\n\n", name)

	// Example 2: Reading numbers
	fmt.Println("Enter your age: ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)
	age, err := strconv.Atoi(ageInput)
	if err != nil {
		fmt.Println("Please enter a valid number")
	} else {
		fmt.Printf("You are %d years old\n\n", age)
	}

	// Example 3: Reading multiple values
	fmt.Println("Enter two numbers separated by space: ")
	numbersInput, _ := reader.ReadString('\n')
	numbers := strings.Fields(strings.TrimSpace(numbersInput))
	if len(numbers) >= 2 {
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		fmt.Printf("Sum of numbers: %d\n\n", num1+num2)
	}

	// Example 4: Reading floating-point numbers
	fmt.Println("Enter your height in meters: ")
	heightInput, _ := reader.ReadString('\n')
	height, err := strconv.ParseFloat(strings.TrimSpace(heightInput), 64)
	if err != nil {
		fmt.Println("Please enter a valid floating-point number")
	} else {
		fmt.Printf("Your height is %.2f meters\n\n", height)
	}

	// Example 5: Reading yes/no input
	fmt.Println("Do you like programming? (yes/no): ")
	answer, _ := reader.ReadString('\n')
	answer = strings.ToLower(strings.TrimSpace(answer))
	if answer == "yes" || answer == "y" {
		fmt.Println("That's great!")
	} else if answer == "no" || answer == "n" {
		fmt.Println("Maybe you'll like it eventually!")
	} else {
		fmt.Println("Invalid input")
	}

	// Example 6: Reading with Scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nEnter some words (one per line, type 'done' to finish):")
	var words []string
	for scanner.Scan() {
		text := scanner.Text()
		if text == "done" {
			break
		}
		words = append(words, text)
	}
	fmt.Printf("You entered %d words: %v\n", len(words), words)
}
