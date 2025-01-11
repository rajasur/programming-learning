package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter you first user-input: ")
	reader := bufio.NewReader(os.Stdin)
	// comma ok or error syntax
	input, _ := reader.ReadString('\n')
	fmt.Printf("Your input %s", input)
}
