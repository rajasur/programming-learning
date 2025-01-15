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
	fmt.Println("Your input: ", input)
	fmt.Printf("Type of this value is %T \n", input)
	fmt.Printf("Your input %s", input)
}
